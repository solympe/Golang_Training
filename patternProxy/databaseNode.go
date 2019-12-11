package patternProxy

type DBNode struct {
	dataBase *dataBase
}

func NewDBNode() *DBNode {
	return &DBNode{dataBase: &dataBase{}}
}

func (n *DBNode) SendData(data string) {
	n.dataBase.SendData(data)
}

func (n DBNode) GetData() {
	n.dataBase.GetData()
}
