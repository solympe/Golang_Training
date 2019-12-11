package patternProxy

import "fmt"

type dataBase struct {
	data string
}

type DBFunctions interface {
	GetData()
	SendData(string)
}

func NewDB() *dataBase {
	return &dataBase{data: ""}
}

func (db *dataBase) SendData(data string) {
	db.data = data
	fmt.Println(db)
}

func (db *dataBase) GetData() {
	if len(db.data) == 0 {
		return
	}
	fmt.Println("Now my data is:", db.data)
}
