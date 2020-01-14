package is_isomorphic

import (
	"reflect"
	"strings"
)

type word struct {
	data string
}

// GetData ...
func (w *word) GetData() string {
	return w.data
}

// IsIsomorphic ...
func (w *word) IsIsomorphic(word2 Worder) bool {
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

// NewWorder ...
func NewWorder(data string) Worder {
	return &word{data: data}
}
