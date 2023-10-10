package main

import (
	"errors"
	"fmt"
)

func aplicarFuncaoESomar(slice []int, funcao func(int) int) (int, error) {
	if len(slice) == 0 {
		return 0, errors.New("O slice está vazio")
	}

	soma := 0
	for _, valor := range slice {
		resultado := funcao(valor)
		soma += resultado
	}
	return soma, nil
}

func main() {

	numeros := []int{1, 2, 3, 4, 5}
	funcao := func(x int) int {
		return x * x
	}
	resultado, err := aplicarFuncaoESomar(numeros, funcao)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Soma dos resultados:", resultado)
	}
}
