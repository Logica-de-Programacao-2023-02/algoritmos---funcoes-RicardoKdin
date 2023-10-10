package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {

	palavras := []string{"Ricardo", "souza", "Moraes"}
	maiusculas, err := letramaiuscula(palavras)

	if err != nil {
		fmt.Print("Erro: ", err)
	} else {
		fmt.Println(maiusculas)
	}
}

func letramaiuscula(s []string) (string, error) {

	if len(s) == 0 {
		return "", fmt.Errorf("O slice está vazio")
	}
	var maiuscula []string
	for _, str := range s {
		if len(str) > 0 && unicode.IsUpper(rune(str[0])) {
			maiuscula = append(maiuscula, str)
		}
	}
	return strings.Join(maiuscula, " "), nil
}
