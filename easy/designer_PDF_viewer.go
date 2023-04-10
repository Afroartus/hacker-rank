package easy

func IndexLetras(h []int32, word string) int32 {

	letras := make(map[int32]int32)
	var altura, anchura int32
	anchura = int32(len(word))

	for _, item := range word {
		num := rune(item) - 96
		for i := 0; i <= len(h); i++ {
			if int32(i+1) == num {
				letras[item] = h[i]
			}
		}
		for _, value := range letras {
			if value > altura {
				altura = value
			}
		}
	}
	return anchura * altura
}
