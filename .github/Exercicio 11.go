package main

import (
	"fmt"
	"math"
)

func main() {

	var num int
	fmt.Print("Digite um número: ")
	fmt.Scan(&num)

	primo, err := numeroprimo(num)
	if err != nil {
		fmt.Println("Erro: ", err)
	} else {
		if primo {
			fmt.Println("É um número primo")
		} else {
			fmt.Println("Não é um número primo")
		}
	}
}

func numeroprimo(num int) (bool, error) {

	if num < 0 {
		return false, fmt.Errorf("Número Negativo")
	}
	if num <= 1 {
		return false, nil
	}
	if num <= 3 {
		return true, nil
	}
	if num%2 == 0 || num%3 == 0 {
		return false, nil
	}
	limite := int(math.Sqrt(float64(num))) + 1
	for i := 5; i <= limite; i += 6 {
		if num%i == 0 || num%(i+2) == 0 {
			return false, nil
		}
	}
	return true, nil
}
