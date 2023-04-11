package Day1

import "fmt"

func minMaxSum(arr []int32) {
	var sumMin, sumMax int64
	var arrSum []int64
	for i := 0; i < len(arr); i++ {
		sum := int64(0)
		for j := 0; j < len(arr); j++ {
			if i == j {
				continue
			}
			sum += int64(arr[j])
		}
		arrSum = append(arrSum, sum)
	}
	sumMin = arrSum[0]
	for i := 0; i < len(arrSum); i++ {
		if arrSum[i] > sumMax {
			sumMax = arrSum[i]
		}
		if arrSum[i] < sumMin {
			sumMin = arrSum[i]
		}
	}
	fmt.Printf("%v %v", sumMin, sumMax)
}
