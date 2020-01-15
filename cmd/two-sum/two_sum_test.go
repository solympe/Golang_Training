package test_two_sum

import (
	"reflect"
	"testing"

	nums "github.com/solympe/Golang_Training/pkg/two-sum"
)

type test struct {
	input nums.Nums
	wait  []int
}

var tests = []test{

	{nums.NewNums([]int{2, 7, 11, 15}, 9), []int{0, 1}},
	{nums.NewNums([]int{0, 1, 2, 3, 999, 9}, 9), []int{0, 5}},
	{nums.NewNums([]int{0, 1, 2, 3, 4, 5}, 3), []int{0, 3}},
	{nums.NewNums([]int{0, 11, 12, 343, 4}, 343), []int{0, 3}},
	{nums.NewNums([]int{9, 9, 3}, 12), []int{0, 2}},
}

func TestTwoSum(t *testing.T) {

	for _, pairs := range tests {
		result := pairs.input.TwoSum()
		if !reflect.DeepEqual(result, pairs.wait) {
			t.Error(pairs.input, "Expected:", pairs.wait, "Result:", result, "FAILED")
		} else {
			reflect.DeepEqual(result, pairs.wait)
			t.Log(pairs.input, "Expected:", pairs.wait, "Result:", result, "OK")
		}
	}
}
