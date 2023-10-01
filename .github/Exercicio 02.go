package main

import (
	"fmt"
	"strings"
)

func main() {

	texto := `Ricardo Souza Moraes`
	vogal := quantidadevogais(texto)
	fmt.Print(`A quantidade de vogais é: `, vogal)

}

func quantidadevogais(texto string) int {

	vogais := `AaEeIiOoUu`
	contador := 0

	for _, letras := range texto {
		if strings.ContainsRune(vogais, letras) {
			contador++
		}
	}

	return contador

}
