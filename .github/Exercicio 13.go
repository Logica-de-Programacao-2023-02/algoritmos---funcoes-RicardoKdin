package main

import (
	"fmt"
	"strconv"
)

func main() {

	var n int
	fmt.Print("Digite um número com duas casas: ")
	fmt.Scan(&n)

	soma, err := somadigitos(n)
	if err != nil {
		fmt.Println("Erro: ", err)
	} else {
		fmt.Println("Soma: ", soma)
	}
}

func somadigitos(n int) (int, error) {

	if n < 0 {
		return 0, fmt.Errorf("Número Negativo")
	}
	numero := strconv.Itoa(n)
	soma := 0
	for _, s := range numero {
		digito, err := strconv.Atoi(string(s))
		if err != nil {
			return 0, err
		}
		soma += digito
	}
	return soma, nil
}
