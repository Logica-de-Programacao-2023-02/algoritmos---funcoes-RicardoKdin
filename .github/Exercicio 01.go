package main

import "fmt"

func main() {

	valores := []int{1, 2, 3, 4, 5}
	media := mediaaritmetica(valores)
	fmt.Print(`A média dos números é: `, media)

}

func mediaaritmetica(slice []int) int {

	soma := 0
	for _, valor := range slice {
		soma += valor
	}

	media := soma / len(slice)
	return media
}
