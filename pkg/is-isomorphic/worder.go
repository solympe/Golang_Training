package is_isomorphic

// Worder is a word interface
type Worder interface {
	IsIsomorphic(word2 Worder) bool
	GetData() string
	createMap() []int
}
