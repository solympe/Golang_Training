package main

import "testing"

type testCase struct {
	input    string
	nameIn   string
	waitName string
	waitType string
}

var tests = []testCase{
	{"evil", "Bob", "Bob", "evil"},
	{"kind", "Terminator", "Terminator", "kind"},
}

func TestAndroidFactory(t *testing.T) {
	for _, pairs := range tests {
		newElement := AndroidFactory(pairs.input)
		newElement.giveName(pairs.nameIn)
		typeofElem := newElement.getType()
		nameofElem := newElement.getName()
		if typeofElem == pairs.waitType && nameofElem == pairs.waitName {
			t.Log("Test passed! Waited:", pairs.waitName, pairs.waitType, ", Gave:", nameofElem, typeofElem)
		} else {
			t.Error("Test failed! Waited:", pairs.waitName, pairs.waitType, ", Gave:", nameofElem, typeofElem)
		}
	}
}
