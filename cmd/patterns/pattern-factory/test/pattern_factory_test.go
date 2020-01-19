package test_test

import (
	"testing"

	"github.com/solympe/Golang_Training/pkg/patterns/pattern-factory/factory"
)

type transportType = string

type testCase struct {
	input    transportType
	waitType transportType
}

var tests = []testCase{
	{"car", "car"},
	{"track", "track"},
}

func TestAndroidFactory(t *testing.T) {
	for _, pairs := range tests {
		androidFactory := factory.NewFactory()
		android := androidFactory.Create(pairs.input).Get()
		if android == pairs.waitType {
			t.Log("Test passed! Waited:", pairs.waitType, ", Gave:", android)
		} else {
			t.Error("Test failed! Waited:", pairs.waitType, ", Gave:", android)
		}
	}
}
