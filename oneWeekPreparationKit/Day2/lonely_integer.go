package Day2

func lonelyInteger(a []int32) int32 {

	b := make(map[int32]int)

	for _, value := range a {
		b[value]++
	}

	for i, value := range b {
		if value == 1 {
			return i
		}
	}

	return 0
}
