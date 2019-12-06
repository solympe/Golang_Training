package main

type intStruct struct {
	elements map[int]bool
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
