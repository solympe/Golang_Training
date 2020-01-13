package set_int

type intStruct struct {
	elements map[int]bool
}

// AddElem ...
func (s *intStruct) AddElem(num int) bool {
	if s.elements[num] == true {
		return false
	}
	s.elements[num] = true
	return true
}

// DeleteElem ...
func (s *intStruct) DeleteElem(num int) bool {
	if s.elements[num] == false {
		return false
	}
	delete(s.elements, num)
	return true
}

// CheckElem ...
func (s *intStruct) CheckElem(num int) bool {
	if s.elements[num] == false {
		return false
	}
	return true
}

// NewIntStructExecutor ...
func NewIntStructExecutor(elements map[int]bool) IntStructExecutor {
	return &intStruct{elements: elements}
}
