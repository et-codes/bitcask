package bitcask

import (
	"fmt"
	"log"
	"os"
)

type DiskStore struct {
	ActiveFile *os.File               // file being appended
	Position   int                    // current offset in the file
	KeyDir     map[string]KeyDirEntry // key directory
}

func NewDiskStore(filename string) Bitcask {
	ds := &DiskStore{
		KeyDir: NewKeyDir(),
	}

	if fileExists(filename) {
		// TODO: handle reopening existing database.
		panic("existing file handling not implemented")
	} else {
		file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("unable to create data file: %v\n", err)
		}
		ds.ActiveFile = file
	}

	return ds
}

func (d *DiskStore) Get(key string) (string, error) {
	entry, found := d.KeyDir[key]
	if !found {
		return "", fmt.Errorf("key %q not found", key)
	}

	value := make([]byte, entry.ValueSize)
	n, err := d.ActiveFile.ReadAt(value, int64(entry.ValuePosition))
	if err != nil || n != int(entry.ValueSize) {
		return "", fmt.Errorf("error reading [%d/%d] bytes from disk: %v",
			n, int(entry.ValueSize), err)
	}

	return string(value), nil
}

// Put stores a new key and value and returns the old value if the key
// already exists, otherwise an empty string.
func (d *DiskStore) Put(key, value string) (string, error) {
	var err error
	var old string

	// If key already exists, get old value
	_, found := d.KeyDir[key]
	if found {
		old, err = d.Get(key)
		if err != nil {
			return "", err
		}
	}

	// Encode the KV
	kv := NewKeyValue(key, value)
	encoded := encodeKV(kv)

	// Write it to disk
	n, err := d.ActiveFile.Write(encoded)
	if err != nil || n != len(encoded) {
		return "", fmt.Errorf("error writing to disk: %v", err)
	}
	d.Position += n

	if err = d.ActiveFile.Sync(); err != nil {
		return "", fmt.Errorf("error syncing to disk: %v", err)
	}

	// Update the KeyDir
	d.KeyDir[key] = KeyDirEntry{
		ValueSize:     kv.ValueSize,
		ValuePosition: uint32(d.Position) - kv.ValueSize,
		Timestamp:     kv.Timestamp,
	}

	return old, nil
}

func (d *DiskStore) Delete(string) (string, error) {
	return "", nil
}

func (d *DiskStore) ListKeys() []string {
	return nil
}

func (d *DiskStore) Close() error {
	return nil
}

func fileExists(filename string) bool {
	if _, err := os.Stat(filename); err != nil {
		return false
	}
	return true
}