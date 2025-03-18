# LRUCache

A thread-safe LRU (Least Recently Used) cache implementation in Go.

## Features

- Thread-safe LRU cache implementation
- Configurable cache capacity
- Basic operations: Get/Set/Remove
- Automatic eviction of least recently used items

## Installation

```bash
go get github.com/hacker4257/lrucache
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/hacker4257/lrucache"
)

func main() {
    // Create a new cache with capacity of 1000
    cache := lrucache.NewLruCache(1000)
    
    // Set a cache entry
    cache.Set("key1", "value1")
    
    // Get a cache entry
    value, exists := cache.Get("key1")
    if exists {
        fmt.Printf("Found value: %v\n", value)
    }
}
```

## API Documentation

### New(capacity int) *Cache
Creates a new LRU cache instance with the specified capacity.

### Get(key interface{}) (interface{}, bool)
Retrieves the value for a given key. Returns false if the key doesn't exist.

### Set(key, value interface{})
Sets a key-value pair in the cache.

### Remove(key interface{})
Removes an entry from the cache by its key.

## Performance

- Time Complexity: O(1) for Get/Set/Remove operations
- Space Complexity: O(capacity) where capacity is the cache size

## Contributing

Issues and Pull Requests are welcome!

## License

MIT License
