package bitcask_test

import (
	"testing"

	"github.com/et-codes/bitcask"
	"github.com/stretchr/testify/assert"
)

func TestBitcask(t *testing.T) {
	t.Run("test put and get", func(t *testing.T) {
		key := "name"
		value := "John Doe"
		b := setUp(t, key, value)

		actual, err := b.Get(key)
		assert.NoError(t, err)
		assert.Equal(t, value, actual)

		newValue := "Jane Doe"
		oldVal, err := b.Put(key, newValue)
		assert.NoError(t, err)
		assert.Equal(t, value, oldVal)
	})

	t.Run("test delete", func(t *testing.T) {
		key := "name"
		value := "John Doe"
		b := setUp(t, key, value)

		actual, err := b.Delete(key)
		assert.NoError(t, err)
		assert.Equal(t, value, actual)

		actual, err = b.Get(key)
		assert.Error(t, err)
		assert.Equal(t, "", actual)
	})

	t.Run("test list keys", func(t *testing.T) {
		b := bitcask.NewMemoryStore()
		keys := []string{"firstName", "lastName", "SSN", "Mom's name"}
		for _, key := range keys {
			_, _ = b.Put(key, "")
		}

		actual := b.ListKeys()
		assert.ElementsMatch(t, keys, actual)
	})
}

func setUp(t *testing.T, key, value string) bitcask.Bitcask {
	t.Helper()
	b := bitcask.NewMemoryStore()
	val, err := b.Put(key, value)
	assert.NoError(t, err)
	assert.Equal(t, "", val)

	return b
}
