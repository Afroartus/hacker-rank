package easy

import "testing"

func TestSolveMeFirst(t *testing.T) {

	var a, b uint32 = 1, 2
	expected := a + b
	got := solveMeFirst(a, b)

	if got != expected {
		t.Errorf("Expected: %v, got: %v", expected, got)
	}
}
