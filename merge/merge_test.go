package main

import (
	"reflect"
	"testing"
)

type testStruct struct {
	input1 []int
	input2 []int
	m int
	n int
	wait []int
}

func TestMerge (t *testing.T) {

	var tests = []testStruct{
		{[]int{1,2,3,0,0,0}, []int{4,5,6}, 3, 3, []int{1,2,3,4,5,6}},
		{[]int{1}, []int{}, 1, 0, []int{1}},
	}

	for _, pairs := range tests {
		result := Merge(pairs.input1, pairs.m, pairs.input2, pairs.n)
		if reflect.DeepEqual(result, pairs.wait) {
			t.Log("Test passed", pairs.wait, "expected", result, "are returned")
		} else {
			t.Error("Error in test! ", pairs.wait, "expected,", result, "are returned")
		}
	}
}