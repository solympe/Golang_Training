package main

import "fmt"

type intStruct struct {
	elements map[int]bool
}

func checkStatus ( m map[int]bool, status bool, operation string) {

	if status == true {
		fmt.Println(operation,"is done!:", m)
	} else {
		fmt.Println(operation, "error!", m)
	}
}

func main() {
	//добавить удалить проверить

	sInt := intStruct{
		elements:  make(map[int]bool),
	}

	checkStatus(sInt.elements, sInt.addElem(10), "Add")

	checkStatus(sInt.elements, sInt.addElem(10), "Add")

	checkStatus(sInt.elements, sInt.addElem(20), "Add")

	checkStatus(sInt.elements, sInt.checkElem(10), "Check")

	checkStatus(sInt.elements, sInt.deleteElem(10), "Delete")

	checkStatus(sInt.elements, sInt.checkElem(10), "Check")

}

func (s *intStruct)addElem(num int) bool{
	if s.elements[num] == true {
		return false
	}
	s.elements[num] = true
	return true
}

func (s *intStruct)deleteElem(num int) bool{
	if s.elements[num] == false {
		return false
	}
	delete(s.elements, num)
	return true
}

func (s *intStruct)checkElem(num int) bool{
	if s.elements[num] == false {
		return false
	}
	return true
}

