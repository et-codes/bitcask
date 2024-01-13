# bitcask

My Go implementation of a persistent key-value store using a variation on the Bitcask specification. Non-persistent in-memory storage can also be used by using filename *:memory:*.

Bitcask specification details: https://riak.com/assets/bitcask-intro.pdf

### Example usage:

```Go
package main

import (
    "fmt"

    "github.com/et-codes/bitcask"
)

func main() {
    b := bitcask.Open("bitcask.db")
    defer b.Close()

    b.Put("firstName", "John")
    b.Put("lastName", "Doe")
    old, _ := b.Put("firstName", "Gary")
    fmt.Println(old) // "John"

    firstName, _ := b.Get("firstName")
    fmt.Println(firstName) // "Gary"

    lastName, _ := b.Delete("lastName")
    fmt.Println(lastName) // "Doe"

    _, err := b.Get("lastName")
    fmt.Println(err) // key "lastName" not found
}
```
