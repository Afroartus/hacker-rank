package easy

import "testing"

func TestBreakingTheRocks(t *testing.T) {

	//OneExercice := []int32{10, 5, 20, 20, 4, 5, 2, 25, 1}
	TwoExercice := []int32{3, 4, 21, 36, 10, 28, 35, 5, 24, 42}
	expect := []int32{2, 4}
	got := breakingTheRecords(TwoExercice)

	for i := 0; i > len(got); i++ {
		if got[i] != expect[i] {
			t.Errorf("Got: %v, Expected: %v", got, expect)
		}
	}
}
