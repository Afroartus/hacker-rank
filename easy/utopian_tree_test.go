package easy

import "testing"

func TestUtopianTree(t *testing.T) {
	arr := []int32{6, 5, 4}
	expected := []int32{15, 14, 7}

	for i := 0; i < len(arr); i++ {
		got := utopianTree(arr[i])
		if got != expected[i] {
			t.Errorf("Expected: %v, got: %v", expected[i], got)
		}
	}

}
