package Day1

import "fmt"

func plusMinus(arr []int32) {
	var positive, negative, zero float64
	long := float64(len(arr))
	result := map[int32]float64{0: 0, 1: 0, 2: 0}

	for _, item := range arr {
		switch {
		case item > 0:
			positive += 1.0
			result[0] = positive
		case item < 0:
			negative += 1.0
			result[1] = negative
		case item == 0:
			zero += 1.0
			result[2] = zero
		}
	}
	for i := 0; i < len(result); i++ {
		fmt.Printf("%.6f\n", result[int32(i)]/long)
	}
}
