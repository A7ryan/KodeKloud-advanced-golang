package main

import (
	"testing"
)

func TestCheckEven(t *testing.T) {
	i := 10
	expected := "Yes"
	res := checkEven(i)
	if expected != res {
		t.Errorf("Expected: %v, Got: %v", expected, res)
	}
}
