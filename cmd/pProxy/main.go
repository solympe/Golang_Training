package main

import (
	"fmt"

	c "github.com/solympe/Golang_Training/pkg/pProxy/cache"
	db "github.com/solympe/Golang_Training/pkg/pProxy/dataBase"
	dbn "github.com/solympe/Golang_Training/pkg/pProxy/dbNode"
)

func main() {
	var cache = c.NewCache()
	var dataBaseOrig = db.NewDB(cache)
	var dataBaseNode = dbn.NewDBNode(cache, dataBaseOrig)

	dataBaseOrig.SendData("new data")
	fmt.Println(dataBaseNode.GetData(), "=", dataBaseOrig.GetData())

	dataBaseNode.SendData("new data 2")
	fmt.Println(dataBaseNode.GetData(), "=", dataBaseOrig.GetData())
}
