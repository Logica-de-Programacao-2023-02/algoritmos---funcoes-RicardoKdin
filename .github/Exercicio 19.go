package main

import (
	"errors"
	"fmt"
)

func ehPrimo(numero int) bool {
	if numero <= 1 {
		return false
	}
	if numero <= 3 {
		return true
	}
	if numero%2 == 0 || numero%3 == 0 {
		return false
	}

	limite := int(numero / 2)
	for i := 5; i <= limite; i += 6 {
		if numero%i == 0 || numero%(i+2) == 0 {
			return false
		}
	}
	return true
}

func primosMenoresOuIguais(numero int) ([]int, error) {
	if numero < 0 {
		return nil, errors.New("Número negativo não é suportado")
	}
	var primos []int
	for i := 2; i <= numero; i++ {
		if ehPrimo(i) {
			primos = append(primos, i)
		}
	}
	return primos, nil
}

func main() {

	numero := 20
	primos, err := primosMenoresOuIguais(numero)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Números primos menores ou iguais a", numero, ":", primos)
	}
}
