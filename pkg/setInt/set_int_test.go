package main

import (
	"testing"
)

type testingS struct {
	num       int
	out       bool
	operation string
	operationCode int
}

var tests = []testingS{

	{10, true,"Adding", 1},
	{20, false,"Deleting/broken", 2},
	{20, true,"Adding", 1},
	{10, true,"Deleting", 2},
	{10, false,"Deleting", 2},
	{20, true,"Checking", 3},

}

func TestSetInt(t *testing.T) {
	sInt := intStruct {
	elements:  make(map[int]bool),
	}

	for _, pairs := range tests {
		var result bool
		switch pairs.operationCode {
		case 1: result = sInt.addElem(pairs.num)
		case 2: result = sInt.deleteElem(pairs.num)
		case 3: result = sInt.checkElem(pairs.num)
		}
		if result == pairs.out {
			t.Log("Success!",  pairs.operation, pairs.num, "was correct in", sInt.elements)
			continue
		} else {
			t.Error("Error!", pairs.operation, pairs.num, "was incorrect in", sInt.elements)
		}
	}
}
