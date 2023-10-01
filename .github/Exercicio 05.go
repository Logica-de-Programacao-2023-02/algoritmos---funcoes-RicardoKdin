package main

import "fmt"

func main() {

	slice := []int{10, 11, 12, 13, 14}
	x := 13
	posicao := posicaoelemento(slice, x)

	fmt.Print(`A posição do valor é: `, posicao)

}

func posicaoelemento(slice []int, x int) int {

	for i, elemento := range slice {
		if elemento == x {
			return i
		}
	}
	return -1
}
