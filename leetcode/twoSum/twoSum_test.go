package main

import (
	"reflect"
	"testing"
)

type test struct {
	input []int
	wait []int
	target int
}

var tests = []test{
	{[]int{2, 7, 11, 15}, []int{0, 1}, 9},
	{[]int{0, 1, 2, 3, 999, 9}, []int{0, 5}, 9},
	{[]int{0, 1, 2, 3, 4, 5}, []int{0, 3}, 3},
	{[]int{0, 11, 12, 343, 4}, []int{0, 3}, 343},
	{[]int{9, 9, 3}, []int{1, 2}, 12}, //broken answer

}

func TestTwoSum (t *testing.T) {

	for _, pairs := range tests {
		result := twoSum(pairs.input, pairs.target)
		if !reflect.DeepEqual(result, pairs.wait) {
			t.Error("Error in", pairs.input, "Im waiting", pairs.wait, "but i have only", result)
		}
		// else does not work wtf.
		if reflect.DeepEqual(result, pairs.wait){
			t.Log(pairs.input, "is correct! Expected:", pairs.wait, "Result:", result)
		}
	}
}