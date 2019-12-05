package main

import (
	"fmt"
	"reflect"

	//"reflect"
	"strings"
)

func main() {

	s := "ala"
	t := "vaa"

	answer := isIsomorphic(s, t)

	fmt.Println("ans: ",answer)
}

func isIsomorphic(s string, t string) bool {
	m1 := createMap(s, t)
	m2 := createMap(t, s)

	fmt.Println(len(m1), len(m2))

	if len(m1) != len(m2) {
		return false
	}
	return true
}

func createMap (s string, t string) map[string]string{
	m := map[string]string{}
	one := strings.Split(s, "")
	two := strings.Split(t, "")
	for i := 0; i < len(one); i++ {
		ok := m[one[i]]
		if ok == "" {
			m[one[i]] = two[i]
		}
	}
	return m
}