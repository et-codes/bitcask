package bitcask

import (
	"encoding/binary"
	"hash/crc32"
	"time"
)

const (
	DELETE     uint8 = 255 // set Tombstone byte to delete
	KEEP       uint8 = 127 // set Tombstone byte to keep
	headerSize       = 17  // CRC, tombstone, timestamp, key size, value size
)

type KeyValue struct {
	CRC       uint32 // CRC check value
	Tombstone uint8  // 255 = delete, 127 = keep
	Timestamp uint32 // Unix timestamp
	KeySize   uint32 // length of the key field
	ValueSize uint32 // length of the value field
	Key       string // key
	Value     string // value
}

func NewKeyValue(key, value string) KeyValue {
	return KeyValue{
		Tombstone: KEEP,
		Timestamp: uint32(time.Now().Unix()),
		KeySize:   uint32(len(key)),
		ValueSize: uint32(len(value)),
		Key:       key,
		Value:     value,
	}
}

// encodeKV encodes a KeyValue struct into a byte slice.
func encodeKV(kv KeyValue) []byte {
	keyLength := len(kv.Key)
	valueLength := len(kv.Value)
	payload := make([]byte, headerSize+keyLength+valueLength)
	payload[4] = byte(kv.Tombstone)                             // 4
	binary.LittleEndian.PutUint32(payload[5:9], kv.Timestamp)   // 5-8
	binary.LittleEndian.PutUint32(payload[9:13], kv.KeySize)    // 9-12
	binary.LittleEndian.PutUint32(payload[13:17], kv.ValueSize) // 13-16
	copy(payload[17:17+keyLength], []byte(kv.Key))              // 17+
	copy(payload[17+keyLength:], []byte(kv.Value))

	// checksum of payload after tombstone
	crc := crc32.ChecksumIEEE(payload[5:])
	binary.LittleEndian.PutUint32(payload[0:4], crc) // 0-3

	return payload
}

// decodeKV turns a slice of bytes into a KeyValue struct.
func decodeKV(data []byte) KeyValue {
	keySize := binary.LittleEndian.Uint32(data[9:13])
	return KeyValue{
		CRC:       binary.LittleEndian.Uint32(data[0:4]),
		Tombstone: uint8(data[4]),
		Timestamp: binary.LittleEndian.Uint32(data[5:9]),
		KeySize:   keySize,
		ValueSize: binary.LittleEndian.Uint32(data[13:17]),
		Key:       string(data[17 : 17+keySize]),
		Value:     string(data[17+keySize:]),
	}
}

// isValidKV returns whether the checksum matches the data.
func isValidKV(data []byte) bool {
	savedCRC := binary.LittleEndian.Uint32(data[0:4])
	actualCRC := crc32.ChecksumIEEE(data[5:])
	return savedCRC == actualCRC
}
