package Day1

import (
	"sort"
)

func findMedian(arr []int32) int32 {
	var arr2 []int

	for _, item := range arr {
		arr2 = append(arr2, int(item))
	}
	sort.Ints(arr2)
	index := len(arr2) / 2
	if len(arr2)%2 == 0 {
		return int32(arr2[index-1])
	}
	return int32(arr2[index])
}
