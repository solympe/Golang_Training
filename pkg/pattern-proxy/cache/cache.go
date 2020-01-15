package cache

// Cache ...
type Сache interface {
	GetData() string
	Send(data string)
}

type cache struct {
	data string
}

// GetData returns data from cache
func (c *cache) GetData() string {
	return c.data
}

// Send updates data in cache
func (c *cache) Send(data string) {
	c.data = data
}

// NewCache returns new instance of cache
func NewCache() Сache {
	return &cache{}
}
