package dataBase

import (
	df "github.com/solympe/Golang_Training/patternProxy/DBFunctions"
)

type dataBase struct {
	data  string
	cache df.DBFunctions
}

// SendData updates data in the main dataBase and cache
func (db *dataBase) SendData(data string) {
	db.data = data
	df.DBFunctions.SendData(db.cache, data)
}

// GetData returns data from main dataBase
func (db dataBase) GetData() string {
	return db.data
}

// NewDB returns new instance of main dataBase
func NewDB(cache df.DBFunctions) df.DBFunctions {
	return &dataBase{cache: cache}
}
