package word

import (
	"reflect"
	"strings"
)

// Word is a word interface
type Word interface {
	IsIsomorphic(word2 Word) bool
	GetData() string
	createMap() []int
}

type word struct {
	data string
}

// GetData ...
func (w *word) GetData() string {
	return w.data
}

// IsIsomorphic ...
func (w *word) IsIsomorphic(word2 Word) bool {
	m1 := w.createMap()
	m2 := word2.createMap()
	return reflect.DeepEqual(m1, m2)
}

func (w *word) createMap() []int {
	m := map[string]int{}
	one := strings.Split(w.data, "")
	intSlice := []int{}
	for i := 0; i < len(one); i++ {
		if m[one[i]] == 0 {
			m[one[i]] = i + 1
			intSlice = append(intSlice, m[one[i]])
			continue
		}
		if m[one[i]] != 0 {
			intSlice = append(intSlice, m[one[i]])
		}
	}
	return intSlice
}

// NewWord ...
func NewWord(data string) Word {
	return &word{data: data}
}
