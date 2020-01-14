package pattern_factor_test

import (
	"testing"

	f "github.com/solympe/Golang_Training/pkg/pattern-factory/factory"
)

type testCase struct {
	input    string
	waitType string
}

var tests = []testCase{
	{"evil", "evil"},
	{"kind", "kind"},
}

func TestAndroidFactory(t *testing.T) {
	for _, pairs := range tests {
		factory := f.NewFactoryCreator()
		typeofElem := factory.CreateAndroid(pairs.input).GetType()
		if typeofElem == pairs.waitType {
			t.Log("Test passed! Waited:", pairs.waitType, ", Gave:", typeofElem)
		} else {
			t.Error("Test failed! Waited:", pairs.waitType, ", Gave:", typeofElem)
		}
	}
}
