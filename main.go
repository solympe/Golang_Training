package main

import (
	db "github.com/solympe/Golang_Training/patternProxy/dataBase"
	dbn "github.com/solympe/Golang_Training/patternProxy/dbNode"
)

func main() {

	var dataBaseOrig db.DBFunctions = db.NewDB("some data 1")
	var dataBaseNode = dbn.NewDBNode(dataBaseOrig)

	dataBaseOrig.GetData()

	dataBaseNode.SendData("some data 2")
	dataBaseOrig.GetData()

	dataBaseNode.SendData("some data 3")
	dataBaseOrig.GetData()
}
