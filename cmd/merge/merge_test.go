package merge_test

import (
	"reflect"
	"testing"

	m "github.com/solympe/Golang_Training/pkg/merge"
)

type testStruct struct {
	input1 m.DataMerger
	input2 m.DataMerger
	wait   []int
}

var tests = []testStruct{
	{m.NewDataMerger([]int{1, 2, 3, 0, 0, 0}, 3), m.NewDataMerger([]int{4, 5, 6}, 3), []int{1, 2, 3, 4, 5, 6}},
	{m.NewDataMerger([]int{1}, 1), m.NewDataMerger([]int{}, 0), []int{1}},
}

func TestMerge(t *testing.T) {
	for _, pairs := range tests {
		result := m.Merge(pairs.input1, pairs.input2)
		if reflect.DeepEqual(result, pairs.wait) {
			t.Log("Test passed", pairs.wait, "expected", result, "are returned")
		} else {
			t.Error("Error in test! ", pairs.wait, "expected,", result, "are returned")
		}
	}
}
