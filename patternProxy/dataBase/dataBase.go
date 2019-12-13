package DataBase

// common interface
type DBFunctions interface {
	GetData() string
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
func (db *dataBase) GetData() string{
	if len(db.data) == 0 {
		return ""
	} else {
		return db.data
	}
}

// data base constructor
func NewDB(data string) DBFunctions {
	return &dataBase{data: data}
}
