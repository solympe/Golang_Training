package dbNode

import (
	df "github.com/solympe/Golang_Training/patternProxy/DBFunctions"
)

type dbNode struct {
	cache    df.DBFunctions
	dataBase df.DBFunctions
}

// SendData updates data in the main dataBase and cache
func (n *dbNode) SendData(data string) {
	df.DBFunctions.SendData(n.cache, data)
	df.DBFunctions.SendData(n.dataBase, data)
}

// GetData returns 'fresh' data from cache
func (n dbNode) GetData() string {
	freshData := df.DBFunctions.GetData(n.cache)
	return freshData
}

// NewDBNode returns new instance of dataBaseNode(proxy)
func NewDBNode(cache df.DBFunctions, db df.DBFunctions) df.DBFunctions {
	return &dbNode{cache, db}
}
