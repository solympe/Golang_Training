package db_functions

// DBFunction is the common interface for data-base, db-node and cache
type DBFunctions interface {
	GetData() string
	SendData(data string)
}
