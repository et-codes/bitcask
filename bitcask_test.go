package bitcask_test

import (
	"os"
	"testing"

	"github.com/et-codes/bitcask"
	"github.com/stretchr/testify/assert"
)

const filename = "test.db"

func TestBitcask(t *testing.T) {
	t.Run("test put and get", func(t *testing.T) {
		b := bitcask.Open(filename)
		defer os.Remove(filename)

		// put works
		key := "name"
		value := "John Doe"
		val, err := b.Put(key, value)
		assert.NoError(t, err)
		assert.Equal(t, "", val)

		// get works
		actual, err := b.Get(key)
		assert.NoError(t, err)
		assert.Equal(t, value, actual)

		// get fails with non-existent key
		_, err = b.Get("bad key")
		assert.Error(t, err)

		// putting same key returns old value
		newValue := "Jane Doe"
		oldVal, err := b.Put(key, newValue)
		assert.NoError(t, err)
		assert.Equal(t, value, oldVal)
		actual, err = b.Get(key)
		assert.NoError(t, err)
		assert.Equal(t, newValue, actual)
	})

	t.Run("test delete", func(t *testing.T) {
		b := bitcask.Open(filename)
		defer os.Remove(filename)

		key := "name"
		value := "John Doe"
		b.Put(key, value)

		actual, err := b.Delete(key)
		assert.NoError(t, err)
		assert.Equal(t, value, actual)

		actual, err = b.Get(key)
		assert.Error(t, err)
		assert.Equal(t, "", actual)
	})

	t.Run("test list keys", func(t *testing.T) {
		b := bitcask.Open(filename)
		defer os.Remove(filename)

		keys := []string{"firstName", "lastName", "SSN", "Mom's name"}
		for _, key := range keys {
			_, _ = b.Put(key, "")
		}

		actual := b.ListKeys()
		assert.ElementsMatch(t, keys, actual)
	})
}

func TestPersistence(t *testing.T) {
	pairs := map[string]string{
		"one":   "I",
		"two":   "II",
		"three": "III",
		"four":  "IV",
		"five":  "V",
	}

	t.Run("loads saved data", func(t *testing.T) {
		b := bitcask.Open(filename)
		defer os.Remove(filename)

		for k, v := range pairs {
			_, err := b.Put(k, v)
			assert.NoError(t, err)
		}

		err := b.Close()
		assert.NoError(t, err)

		b = bitcask.Open(filename)
		for k, v := range pairs {
			val, err := b.Get(k)
			assert.NoError(t, err)
			assert.Equal(t, v, val)
		}
	})

	t.Run("doesn't load deleted data", func(t *testing.T) {
		b := bitcask.Open(filename)
		defer os.Remove(filename)

		for k, v := range pairs {
			_, err := b.Put(k, v)
			assert.NoError(t, err)
		}

		b.Delete("five")

		err := b.Close()
		assert.NoError(t, err)

		b = bitcask.Open(filename)
		for k, v := range pairs {
			val, err := b.Get(k)
			if k == "five" {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, v, val)
			}
		}
	})
}
