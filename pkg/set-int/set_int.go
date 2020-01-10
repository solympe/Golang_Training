package set_int

// IntManager ...
type IntManager interface {
	AddElem(num int) bool
	DeleteElem(num int) bool
	CheckElem(num int) bool
}

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

// NewIntStruct ...
func NewIntStruct(elements map[int]bool) IntManager {
	return &intStruct{elements: elements}
}
