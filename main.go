package main

import (
	db "github.com/solympe/Golang_Training/patternProxy/dataBase"
	dbn "github.com/solympe/Golang_Training/patternProxy/dbNode"
)

func main() {

	var dataBaseOrig = db.NewDB("some data 1")
	var dataBaseNode = dbn.NewDBNode(dataBaseOrig)
	dataBaseNode.GetData()

	dataBaseNode.SendData("some data 2")
	dataBaseNode.GetData()

	dataBaseNode.SendData("some data 3")
	dataBaseNode.GetData()
}
