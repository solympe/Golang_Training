package data_base

import (
	"github.com/solympe/Golang_Training/pkg/pattern-proxy/db-functions"
)

type dataBase struct {
	data  string
	cache db_functions.DBFunctions
}

// SendData updates data in the main data-base and cache
func (db *dataBase) SendData(data string) {
	db.data = data
	db_functions.DBFunctions.SendData(db.cache, data)
}

// GetData returns data from main data-base
func (db dataBase) GetData() string {
	return db.data
}

// NewDB returns new instance of main data-base
func NewDB(cache db_functions.DBFunctions) db_functions.DBFunctions {
	return &dataBase{cache: cache}
}
