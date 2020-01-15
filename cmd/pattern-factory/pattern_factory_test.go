package pattern_factor_test

import (
	"testing"

	"github.com/solympe/Golang_Training/pkg/pattern-factory/factory"
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
		androidFactory := factory.NewFactory()
		android := androidFactory.CreateAndroid(pairs.input).GetType()
		if android == pairs.waitType {
			t.Log("Test passed! Waited:", pairs.waitType, ", Gave:", android)
		} else {
			t.Error("Test failed! Waited:", pairs.waitType, ", Gave:", android)
		}
	}
}
