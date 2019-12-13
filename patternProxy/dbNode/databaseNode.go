package dbNode

import db "github.com/solympe/Golang_Training/patternProxy/dataBase"

// data base node struct (proxy)
type dbNode struct {
	dataBase db.DBFunctions
	cache string
}

// sending data to proxy
func (n *dbNode) SendData(data string) {
	n.dataBase.SendData(data)    //update dataBase info
	n.cache = data               //update cache info
}

// getting data from proxy
func (n *dbNode) GetData() string{
	return "Now my data is:" + n.cache
}

// proxy constructor
func NewDBNode(datab db.DBFunctions) db.DBFunctions {
	return &dbNode{datab, datab.GetData()}
}
