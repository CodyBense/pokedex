package pokecache

type Cache struct {
	cache map[string][]byte
}

// NewCache ...
func NewCache() Cache {
	return Cache{
		cache: make(map[string][]byte),
	}
}

// Add ...
func (c *Cache) Add(key string, value []byte) {
	c.cache[key] = value
}

// Get ...
func (c *Cache) Get(key string) ([]byte, bool) {
	value, ok := c.cache[key]
	return value, ok
}
