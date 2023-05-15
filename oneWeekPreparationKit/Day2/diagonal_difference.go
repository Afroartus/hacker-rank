package Day2

import "math"

func diagonalDifference(arr [][]int32) int32 {
	cont := 1
	var a int32
	var b int32
	for i := 0; i < len(arr); i++ {
		a += arr[i][i]
		b += arr[i][len(arr)-cont]
		cont++
	}
	return int32(math.Abs(float64(a - b)))
}
