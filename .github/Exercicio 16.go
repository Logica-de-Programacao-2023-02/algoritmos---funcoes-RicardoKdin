package main

import (
	"fmt"
	"strings"
)

func main() {

	s := "Ricardo Souza Moraes"
	texto, err := subvogais(s)

	if err != nil {
		fmt.Print("Erro: ", err)
	} else {
		fmt.Println(texto)
	}
}

func subvogais(s string) (string, error) {

	if s == "" {
		return "", fmt.Errorf("A string está vazia")
	}
	vogais := "AEIOUaeiou"
	for _, vogal := range vogais {
		s = strings.ReplaceAll(s, string(vogal), "*")
	}
	return s, nil
}
