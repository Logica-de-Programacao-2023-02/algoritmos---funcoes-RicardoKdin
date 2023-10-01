package main

import (
	"fmt"
	"strings"
)

func main() {

	slice := []string{}
	resultado, erro := concatenacao(slice)

	if erro != nil {
		fmt.Print(erro)
	} else {
		fmt.Println(`Resultado: `, resultado)
	}

}

func concatenacao(slice []string) (string, error) {

	if len(slice) == 0 {
		return ``, fmt.Errorf(`O slice está vazio`)
	}
	resultado := strings.Join(slice, `,`)
	return resultado, nil
}
