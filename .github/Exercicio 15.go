package main

import "fmt"

func main() {

	numeros := []int{1, 2, 3, 4, 5}
	var numero int

	fmt.Print("Escreva um número: ")
	fmt.Scan(&numero)

	presente, err := numeropresente(numero, numeros)

	if err != nil {
		fmt.Println("Erro: ", err)
	} else {
		if presente {
			fmt.Println(numero, "está presente")
		} else {
			fmt.Println(numero, "não está presente")
		}
	}
}

func numeropresente(n int, s []int) (bool, error) {

	if len(s) < 0 {
		return false, fmt.Errorf("Slice vazio")
	}
	for _, num := range s {
		if n == num {
			return true, nil
		}
	}
	return false, nil
}
