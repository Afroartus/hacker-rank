package pruebasRandom

import "fmt"

func pruebas() {
	algo := make(map[int]int)

	for i := 0; i < 10; i++ {
		algo[i] = i
	}

	for key, value := range algo {
		fmt.Println(key, value)
	}

}
