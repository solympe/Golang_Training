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
func IsIsomorphic(s string, t string) bool {
	m1 := createMap(s)
	m2 := createMap(t)
	return reflect.DeepEqual(m1, m2)
}

func createMap(s string) []int {
	m := map[string]int{}
	one := strings.Split(s, "")
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
