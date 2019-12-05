package main

import (
	"reflect"
	"testing"
)

func TestTwoSum (t *testing.T) {
		arr := []int{0, 2, 3, 6, 2, 2, 2, 7, 11, 15, 9}

	var ansr1 = twoSum(arr, 9)
	if !reflect.DeepEqual(ansr1, []int{0, 10}) {
			t.Errorf("Oops!")
		}
}

