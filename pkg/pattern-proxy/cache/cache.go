package cache

// Cache ...
type Cache interface {
	Get() (mySpecificData string)
	Set(mySpecificData string)
}

type cache struct {
	mySpecificData string
}

// Get returns data from cache
func (c *cache) Get() string {
	return c.mySpecificData
}

// Set updates data in cache
func (c *cache) Set(mySpecificData string) {
	c.mySpecificData = mySpecificData
}

// NewCache returns new instance of cache
func NewCache() Cache {
	return &cache{}
}
