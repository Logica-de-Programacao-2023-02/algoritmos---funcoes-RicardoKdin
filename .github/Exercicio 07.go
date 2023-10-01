package main

import (
	"fmt"
)

type FuncaoMapeamento func(int) int

func aplicarFuncaoEmSlice(slice []int, fn FuncaoMapeamento) ([]int, error) {
	if len(slice) == 0 {
		return nil, fmt.Errorf("O slice está vazio")
	}

	resultado := make([]int, len(slice))
	for i, valor := range slice {
		resultado[i] = fn(valor)
	}

	return resultado, nil
}

func main() {

	numeros := []int{1, 2, 3, 4, 5}

	dobrar := func(x int) int {
		return x * 2
	}

	resultado, err := aplicarFuncaoEmSlice(numeros, dobrar)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Resultado após aplicar a função:", resultado)
	}
}
