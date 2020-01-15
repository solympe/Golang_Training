package cache

type Сache interface {
	GetData() string
	SendData(data string)
}

type cache struct {
	data string
}

// GetData returns data from cache
func (c *cache) GetData() string {
	return c.data
}

// SendData updates data in cache
func (c *cache) SendData(data string) {
	c.data = data
}

// NewCache returns new instance of cache
func NewCache() Сache {
	return &cache{}
}
