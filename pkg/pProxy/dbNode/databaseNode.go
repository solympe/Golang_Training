package dbNode

import (
	"github.com/solympe/Golang_Training/pkg/pProxy/DBFunctions"
)

type dbNode struct {
	cache    DBFunctions.DBFunctions
	dataBase DBFunctions.DBFunctions
}

// SendData updates data in the main dataBase and cache
func (n *dbNode) SendData(data string) {
	DBFunctions.DBFunctions.SendData(n.cache, data)
	DBFunctions.DBFunctions.SendData(n.dataBase, data)
}

// GetData returns 'fresh' data from cache
func (n dbNode) GetData() string {
	freshData := DBFunctions.DBFunctions.GetData(n.cache)
	return freshData
}

// NewDBNode returns new instance of dataBaseNode(proxy)
func NewDBNode(cache DBFunctions.DBFunctions, db DBFunctions.DBFunctions) DBFunctions.DBFunctions {
	return &dbNode{cache, db}
}
