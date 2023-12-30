package bitcask

type KeyDirEntry struct {
	FileId        uint32
	ValueSize     uint32
	ValuePosition uint32
	Timestamp     uint32
}

func NewKeyDir() map[string]KeyDirEntry {
	return make(map[string]KeyDirEntry)
}
