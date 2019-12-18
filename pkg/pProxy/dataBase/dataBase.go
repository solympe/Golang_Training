package dataBase

import (
	"github.com/solympe/Golang_Training/pkg/pProxy/DBFunctions"
)

type dataBase struct {
	data  string
	cache DBFunctions.DBFunctions
}

// SendData updates data in the main dataBase and cache
func (db *dataBase) SendData(data string) {
	db.data = data
	DBFunctions.DBFunctions.SendData(db.cache, data)
}

// GetData returns data from main dataBase
func (db dataBase) GetData() string {
	return db.data
}

// NewDB returns new instance of main dataBase
func NewDB(cache DBFunctions.DBFunctions) DBFunctions.DBFunctions {
	return &dataBase{cache: cache}
}
