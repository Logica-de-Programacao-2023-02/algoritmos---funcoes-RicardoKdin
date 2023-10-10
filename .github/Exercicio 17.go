package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

func ordemalfabetica(s []string) (string, error) {
	if len(s) == 0 {
		return "", errors.New("O slice está vazio")
	}

	sort.Strings(s)

	resultado := strings.Join(s, " ")

	return resultado, nil
}

func main() {

	palavras := []string{"Ricardo", "Souza", "Moraes"}
	resultado, err := ordemalfabetica(palavras)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Strings ordenadas:", resultado)
	}
}
