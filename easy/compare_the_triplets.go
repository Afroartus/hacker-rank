package easy

func compareTheTriplets(a, b []int32) []int32 {

	var aPoint, bPoint int32 = 0, 0
	var arr []int32
	for i, aNum := range a {
		bNum := b[i]
		switch {
		case aNum > bNum:
			aPoint++
		case aNum < bNum:
			bPoint++
		}
	}
	arr = append(arr, aPoint, bPoint)
	return arr
}
