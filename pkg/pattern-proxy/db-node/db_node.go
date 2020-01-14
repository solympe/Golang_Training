package db_node

import (
	"github.com/solympe/Golang_Training/pkg/pattern-proxy/db-functions"
)

type dbNode struct {
	cache    db_functions.DBFunctions
	dataBase db_functions.DBFunctions
}

// SendData updates data in the main data-base and cache
func (n *dbNode) SendData(data string) {
	db_functions.DBFunctions.SendData(n.cache, data)
	db_functions.DBFunctions.SendData(n.dataBase, data)
}

// GetData returns 'fresh' data from cache
func (n dbNode) GetData() string {
	freshData := db_functions.DBFunctions.GetData(n.cache)
	return freshData
}

// NewDBNode returns new instance of dataBaseNode(proxy)
func NewDBNode(cache db_functions.DBFunctions, db db_functions.DBFunctions) db_functions.DBFunctions {
	return &dbNode{cache, db}
}
