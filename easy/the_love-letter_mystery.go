package easy

// s := abcd
// 97 a 122

func theLoveLetterMystery(s string) int32 {
	arr := []rune(s)
	i := 0
	j := len(arr) - 1
	var countOperation int32 = 0
	for j >= len(arr)-(len(arr)/2) {
		runeUno := arr[i]
		runeDos := arr[j]
		for runeUno != runeDos {
			if runeUno > runeDos {
				runeUno -= 1
			} else {
				runeDos -= 1
			}
			countOperation += 1
		}
		//fmt.Println(string(runeUno), " ", string(runeDos), " ", countOperation)
		i++
		j--
	}

	return countOperation
}
