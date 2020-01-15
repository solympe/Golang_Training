package database

// DataBase ...
type DataBase interface {
	GetData() string
	SendData(data string)
}

type dataBase struct {
	data  string
	cache DataBase
}

// SendData updates data in the main database and cache
func (db *dataBase) SendData(data string) {
	db.data = data
	DataBase.SendData(db.cache, data)
}

// GetData returns data from main database
func (db dataBase) GetData() string {
	return db.data
}

// NewDataBase returns new instance of main database
func NewDataBase(cache DataBase) DataBase {
	return &dataBase{cache: cache}
}
