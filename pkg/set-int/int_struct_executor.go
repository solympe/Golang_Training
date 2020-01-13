package set_int

// IntStructExecutor...
type IntStructExecutor interface {
	AddElem(num int) bool
	DeleteElem(num int) bool
	CheckElem(num int) bool
}
