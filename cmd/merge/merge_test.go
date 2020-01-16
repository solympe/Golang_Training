package merge_test

import (
	"reflect"
	"testing"

	merge "github.com/solympe/Golang_Training/pkg/merge"
)

type testStruct struct {
	input1 merge.Data
	input2 merge.Data
	wait   []int
}

var tests = []testStruct{
	{merge.NewData([]int{1, 2, 3, 0, 0, 0}, 3), merge.NewData([]int{4, 5, 6}, 3), []int{1, 2, 3, 4, 5, 6}},
	{merge.NewData([]int{1}, 1), merge.NewData([]int{}, 0), []int{1}},
}

func TestMerge(t *testing.T) {
	for _, pairs := range tests {
		result := pairs.input1.Merge(pairs.input2)
		if reflect.DeepEqual(result, pairs.wait) {
			t.Log("Test passed", pairs.wait, "expected", result, "are returned")
		} else {
			t.Error("Error in test! ", pairs.wait, "expected,", result, "are returned")
		}
	}
}
