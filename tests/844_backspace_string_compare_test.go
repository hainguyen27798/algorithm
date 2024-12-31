package tests

import (
	"github.com/hainguyen27798/algorithm/solutions"
	"testing"
)

func TestCase844_1(t *testing.T) {
	if ok := solutions.BackspaceCompare("ab#c", "ad#c"); !ok {
		t.Fatalf(`BackspaceCompare("") = %t; expected = %t`, ok, true)
	}
}

func TestCase844_2(t *testing.T) {
	if ok := solutions.BackspaceCompare("ab##", "c#d#"); !ok {
		t.Fatalf(`BackspaceCompare("") = %t; expected = %t`, ok, true)
	}
}
