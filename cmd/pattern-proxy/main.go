package main

import (
	"fmt"

	"github.com/solympe/Golang_Training/pkg/pattern-proxy/cache"
	"github.com/solympe/Golang_Training/pkg/pattern-proxy/database"
	"github.com/solympe/Golang_Training/pkg/pattern-proxy/proxy"
)

func main() {
	var dataCache = cache.NewCache()
	var dataBase = database.NewDataBase(dataCache)
	var dbProxy = proxy.NewProxy(dataCache, dataBase)

	dataBase.SendData("new data")
	fmt.Println(dbProxy.GetData(), "=", dataBase.GetData())

	dbProxy.SendData("new data 2")
	fmt.Println(dbProxy.GetData(), "=", dataBase.GetData())
}
