package easy

func utopianTree(n int32) int32 {
	treeH := 0
	for i := 0; i <= int(n); i++ {
		switch {
		case i == 0:
			treeH = 1
			continue
		case i == 1:
			treeH *= 2
			continue
		case i%2 == 0:
			treeH += 1
			continue
		case i%1 == 0, i%3 == 0, i%5 == 0, i%7 == 0:
			treeH *= 2
		}
	}
	return int32(treeH)
}
