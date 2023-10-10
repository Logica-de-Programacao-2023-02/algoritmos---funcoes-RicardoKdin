package main

import (
	"fmt"
	"strings"
)

func main() {

	s := "Ricardo Souza Moraes"
	palavra, err := slicestring(s)

	if err != nil {
		fmt.Print("Erro", err)
	} else {
		fmt.Println("Palavras: ", palavra)
	}

}

func slicestring(s string) ([]string, error) {

	if s == "" {
		return nil, fmt.Errorf("A string está vazia")
	}

	palavras := strings.Fields(s)
	return palavras, nil
}
