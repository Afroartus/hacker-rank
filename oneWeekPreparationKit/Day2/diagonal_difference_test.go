package Day2

import (
	"fmt"
	"testing"
)

func TestDiagonalDifference(t *testing.T) {
	//a := [][]int32{{1, 2, 3}, {4, 5, 6}, {9, 8, 9}}
	b := [][]int32{
		{-1, 1, -7, -8},
		{-10, -8, -5, -2},
		{0, 9, 7, -1},
		{4, 4, -2, 1},
	}

	fmt.Println(diagonalDifference(b))
}
