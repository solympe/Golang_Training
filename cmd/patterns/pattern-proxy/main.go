package main

import (
	"fmt"

	"github.com/solympe/Golang_Training/pkg/patterns/pattern-proxy/cache"
	"github.com/solympe/Golang_Training/pkg/patterns/pattern-proxy/database"
	"github.com/solympe/Golang_Training/pkg/patterns/pattern-proxy/proxy"
)

func main() {
	var dataBase = database.NewDataBase()
	var dataCache = cache.NewCache()

	var dbProxy = proxy.NewProxy(dataBase, dataCache)

	dataBase.Set("new data" )
	fmt.Println(dbProxy.Get())
	fmt.Println(dataBase.Get())

	dataBase.Set("new data2" )
	fmt.Println(dbProxy.Get())
	fmt.Println(dataBase.Get())

	dbProxy.Set("data 3")
	fmt.Println(dbProxy.Get())
	fmt.Println(dataBase.Get())
}
