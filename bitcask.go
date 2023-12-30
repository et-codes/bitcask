package bitcask

type Bitcask interface {
	Get(string) (string, error)
	Put(string, string) (string, error)
	Delete(string) (string, error)
	ListKeys() []string
	Close() error
}