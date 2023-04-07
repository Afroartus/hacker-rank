package easy

import "testing"

func TestTheLoveLetterMystery(t *testing.T) {
	arr := []string{"abc", "abcba", "abcd", "cba"}
	expected := []int32{2, 0, 4, 2}
	for i := 0; i < len(arr); i++ {
		got := theLoveLetterMystery(arr[i])
		if got != expected[i] {
			t.Errorf("Expected: %d, got: %d", expected, got)
		}
	}
}
