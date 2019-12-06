package main

import (
	"fmt"
	"strings"
	"reflect"
)

func main() {

	s := "aada"
	t := "daad"

	answer := isIsomorphic(s, t)

	fmt.Println("ans: ",answer)
}

func isIsomorphic(s string, t string) bool {
	m1 := createMap(s)
	m2 := createMap(t)
	return reflect.DeepEqual(m1, m2)
}

func createMap (s string) []int{
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