package easy

import "testing"

func TestBeautifulBinaryString(t *testing.T) {
	arrayNumber := []string{"0101010", "01100", "0100101010", "010"}
	expected := []int32{2, 0, 3, 1}

	for i := 0; i < len(arrayNumber); i++ {
		got := BeautifulBinaryString(arrayNumber[i])

		if got != expected[i] {
			t.Errorf("Expected: %v, got: %v", expected, got)
		}
	}
}
