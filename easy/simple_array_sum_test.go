package easy

import "testing"

func TestSimpleArraySum(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	got := simpleArraySum(arr)
	var expected int32 = 15

	if got != expected {
		t.Errorf("Expected: %v, got: %v", expected, got)
	}
}
