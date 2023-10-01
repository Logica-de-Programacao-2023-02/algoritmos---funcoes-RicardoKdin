package main

import "fmt"

func main() {

	valores := []int{1, 5, 9, 3, 7}
	segundovalor := segundomaior(valores)
	fmt.Print(`O segundo maior valor é: `, segundovalor)

}

func segundomaior(slice []int) int {

	primeirovalor := slice[0]
	segundovalor := slice[1]

	if primeirovalor < segundovalor {
		primeirovalor, segundovalor = segundovalor, primeirovalor
	}

	for i := 2; i < len(slice); i++ {
		if slice[i] > primeirovalor {
			segundovalor = primeirovalor
			primeirovalor = slice[i]
		} else if slice[i] > segundovalor && slice[i] != primeirovalor {
			segundovalor = slice[i]
		}
	}
	return segundovalor
}
