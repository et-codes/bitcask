package bitcask

import (
	"hash/crc32"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeKV(t *testing.T) {
	kv := NewKeyValue("name", "John")
	payload := encodeKV(kv)
	decoded := decodeKV(payload)
	assert.Equal(t, crc32.ChecksumIEEE(payload[5:]), decoded.CRC)
	assert.Equal(t, kv.Tombstone, decoded.Tombstone)
	assert.Equal(t, kv.Timestamp, decoded.Timestamp)
	assert.Equal(t, kv.KeySize, decoded.KeySize)
	assert.Equal(t, kv.ValueSize, decoded.ValueSize)
	assert.Equal(t, kv.Key, decoded.Key)
	assert.Equal(t, kv.Value, decoded.Value)
}
