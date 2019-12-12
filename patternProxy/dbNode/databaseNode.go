package dbNode

import db "github.com/solympe/Golang_Training/patternProxy/dataBase"

type dbNode struct {
	dataBase db.DBFunctions
}

func (n *dbNode) SendData(data string) {
	n.dataBase.SendData(data)
}

func (n dbNode) GetData() {
	n.dataBase.GetData()
}

func NewDBNode(datab db.DBFunctions) db.DBFunctions {
	return &dbNode{datab}
}
