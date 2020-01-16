package database

// DataBase ...
type DataBase interface {
	Get() (mySpecificData string)
	Set(mySpecificData string)
}

type dataBase struct {
	mySpecificData string
}

// Set updates mySpecificData in the main database and cache
func (db *dataBase) Set(mySpecificData string) {
	db.mySpecificData = mySpecificData
}

// Get returns mySpecificData from main database
func (db *dataBase) Get() string {
	return db.mySpecificData
}

// NewDataBase returns new instance of main database
func NewDataBase() DataBase {
	return &dataBase{}
}
