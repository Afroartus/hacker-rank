package easy

import "testing"

func TestCompareTheTriplets(t *testing.T) {
	a := [][]int32{{5, 6, 7}, {2, 8, 4}, {1, 8, 13}}
	b := [][]int32{{2, 6, 1}, {12, 8, 34}, {1, 28, 3}}
	expected := [][]int32{{2, 0}, {0, 2}, {1, 1}}
	var got []int32
	for i := 0; i < len(a); i++ {
		got = compareTheTriplets(a[i], b[i])
		for j := 0; j < len(got); j++ {
			if got[j] != expected[i][j] {
				t.Errorf("Expected: %v, got: %v", expected, got)
			}
		}
	}
}
