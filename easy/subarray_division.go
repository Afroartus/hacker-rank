package easy

func subArrayDivision(s []int32, d int32, m int32) int32 {
	c := 0
	for i := 0; i < len(s); i++ {
		if r := s[i] + s[i+1]; r == d {
			c++
		}
	}

	return int32(c)
}
