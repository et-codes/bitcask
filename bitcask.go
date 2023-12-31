package bitcask

type Bitcask interface {
	Get(key string) (string, error)
	Put(key, value string) (string, error)
	Delete(key string) (string, error)
	ListKeys() []string
	Close() error
}

func New(fileName string) Bitcask {
	if fileName == ":memory:" {
		return NewMemoryStore()
	}
	return NewDiskStore(fileName)
}
