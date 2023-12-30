package bitcask

import (
	"log"
	"os"
)

type DiskStore struct{
	ActiveFile *os.File // file being appended
}

func NewDiskStore(filename string) Bitcask {
	ds := &DiskStore{}
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

func (d *DiskStore) Get(string) (string, error) {
	return "", nil
}

func (d *DiskStore) Put(string, string) (string, error) {
	return "", nil
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