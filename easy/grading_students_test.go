package easy

import "testing"

func TestGradingStudents(t *testing.T) {
	arr := []int32{73, 67, 38, 33}
	expected := []int32{75, 67, 40, 33}
	got := gradingStudents(arr)
	for i, item := range got {
		if item != expected[i] {
			t.Errorf("Expected: %v, got: %v", expected, got)
		}
	}
}
