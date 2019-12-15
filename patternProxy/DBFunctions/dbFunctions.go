package DBFunctions

// DBFunction is the common interface for dataBase, dbNode and cache
type DBFunctions interface {
	GetData() string
	SendData(data string)
}
