package cache

import (
	"github.com/solympe/Golang_Training/pkg/pattern-proxy/db-functions"
)

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
func NewCache() db_functions.DBFunctions {
	return &cache{}
}
