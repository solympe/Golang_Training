package main

import (
	"tasks/patternProxy"
)

func main() {

	var dataBase patternProxy.DBFunctions = patternProxy.NewDB()
	var dbNode patternProxy.DBFunctions = patternProxy.NewDBNode()


	dbNode.SendData("some data")

	dbNode.GetData()
	dataBase.GetData()
}
