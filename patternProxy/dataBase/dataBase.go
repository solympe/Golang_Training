package DataBase

import "fmt"

// common interface
type DBFunctions interface {
	GetData()
	SendData(string)
}

// data base struct (subject)
type dataBase struct {
	data string
}

// sending data to data base
func (db *dataBase) SendData(data string) {
	db.data = data
}

// getting data from data base
func (db *dataBase) GetData() {
	if len(db.data) == 0 {
		return
	}
	fmt.Println("Now my data is:", db.data)
}

// data base constructor
func NewDB(data string) DBFunctions {
	return &dataBase{data: data}
}
