package bitcask

type Bitcask interface {
	Get(key string) (string, error)
	Put(key, value string) (string, error)
	Delete(key string) (string, error)
	ListKeys() []string
	Close() error
}

// Open opens an existing Bitcask database file, or creates a new one
// if the filename does not exist. Use filename ":memory:" to create an
// in-memory database.
func Open(fileName string) Bitcask {
	if fileName == ":memory:" {
		return NewMemoryStore()
	}
	return NewDiskStore(fileName)
}
