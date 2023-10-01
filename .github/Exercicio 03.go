package main

import (
	"fmt"
	"strings"
)

func main() {

	slice := []string{`Ricardo`, `Souza`, `Moraes`}
	resultado := concatenacao(slice)
	fmt.Print(`Resultado: `, resultado)

}

func concatenacao(slice []string) string {

	resultado := strings.Join(slice, ` `)
	return resultado

}
