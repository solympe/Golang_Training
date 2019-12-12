package DataBase

import "fmt"

type DBFunctions interface {
	GetData()
	SendData(string)
}

type DataBase struct {
	data string
}

func (db *DataBase) SendData(data string) {
	db.data = data
}

func (db *DataBase) GetData() {
	if len(db.data) == 0 {
		return
	}
	fmt.Println("Now my data is:", db.data)
}

func NewDB(data string) *DataBase {
	return &DataBase{data: data}
}
