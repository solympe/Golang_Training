package dbNode

import db "github.com/solympe/Golang_Training/patternProxy/dataBase"

// data base node struct (proxy)
type dbNode struct {
	dataBase db.DBFunctions
}

// sending data to proxy
func (n *dbNode) SendData(data string) {
	n.dataBase.SendData(data)
}

// getting data from proxy
func (n *dbNode) GetData() {
	n.dataBase.GetData()
}

// proxy constructor
func NewDBNode(datab db.DBFunctions) db.DBFunctions {
	return &dbNode{datab}
}
