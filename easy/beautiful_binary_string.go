package easy

func BeautifulBinaryString(a string) int32 {

	arr := []rune(a)
	var numberOfOperation int32 = 0
	i := 0
	for i < len(arr)-2 {
		if arr[i : i+3][0] == '0' && arr[i : i+3][1] == '1' && arr[i : i+3][2] == '0' {
			arr[i+1] = '0'
			numberOfOperation++
			i += 3
		} else {
			i++
		}
	}

	return numberOfOperation

}
