package Day2

func countingSort1(arr []int32) []int32 {
	a := make(map[int]int)
	arr2 := make([]int32, 100)
	var b []int32

	for _, value := range arr {
		a[int(value)]++
	}
	for i, value := range a {
		arr2[i] = int32(value)
	}
	for _, value := range arr2 {
		if value != 0 {
			b = append(b, value)
		}
	}

	return arr2

}
