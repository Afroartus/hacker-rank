package easy

func gradingStudents(grades []int32) []int32 {
	var b []int32
	for _, a := range grades {
		if a >= 38 {
			switch a % 10 {
			case 3, 8:
				a += 2
			case 4, 9:
				a += 1
			}
		}
		b = append(b, a)
	}
	return b
}
