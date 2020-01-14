package set_int_test

import (
	"testing"

	si "github.com/solympe/Golang_Training/pkg/set-int"
)

type testingS struct {
	num           int
	out           bool
	operation     string
	operationCode int
}

var tests = []testingS{

	{10, true, "Adding", 1},
	{20, false, "Deleting/broken", 2},
	{20, true, "Adding", 1},
	{10, true, "Deleting", 2},
	{10, false, "Deleting", 2},
	{20, true, "Checking", 3},
}

func TestSetInt(t *testing.T) {
	sInt := si.NewIntStructExecutor(make(map[int]bool))

	for _, pairs := range tests {
		var result bool
		switch pairs.operationCode {
		case 1:
			result = sInt.AddElem(pairs.num)
		case 2:
			result = sInt.DeleteElem(pairs.num)
		case 3:
			result = sInt.CheckElem(pairs.num)
		}
		if result == pairs.out {
			t.Log(pairs.operation, pairs.num, " OK")
			continue
		} else {
			t.Error(pairs.operation, pairs.num, " FAILED")
		}
	}
}
