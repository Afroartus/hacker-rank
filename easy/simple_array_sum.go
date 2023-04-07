package easy

func simpleArraySum(arr []int) int32 {
	var a int
	for _, b := range arr {
		a += b
	}
	return int32(a)
}
