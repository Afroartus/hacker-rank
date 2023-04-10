package easy

import "testing"

func TestName(t *testing.T) {
	h1 := []int32{1, 3, 1, 3, 1, 4, 1, 3, 2, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5}
	word1 := "abc"
	expected1 := int32(9)

	got := IndexLetras(h1, word1)

	if got != expected1 {
		t.Errorf("Expected: %v, got: %v", expected1, got)
	}

	h2 := []int32{1, 3, 1, 3, 1, 4, 1, 3, 2, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 7}
	word2 := "zaba"
	expected2 := int32(28)

	got = IndexLetras(h2, word2)

	if got != expected2 {
		t.Errorf("Expected: %v, got: %v", expected2, got)
	}

}
