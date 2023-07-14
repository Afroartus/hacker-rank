package easy

func breakingTheRecords(scores []int32) []int32 {

	var (
		min            = scores[0]
		max            = scores[0]
		scoreMin int32 = 0
		scoreMax int32 = 0
	)

	for _, score := range scores {
		switch {
		case score > max:
			max = score
			scoreMax++
		case score < min:
			min = score
			scoreMin++
		}
	}

	return []int32{scoreMax, scoreMin}

}
