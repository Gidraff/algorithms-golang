package numeric

import "testing"

func TestReverse(t *testing.T) {
	// Given a string, return its reverse
	inputStr := "cat"
	expectedStr := "tac"
	result := Reverse(inputStr)
	if result != expectedStr {
		t.Errorf("Expected: %s, got: %s", expectedStr, result)
	}
}
