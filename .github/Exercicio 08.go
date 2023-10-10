package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	pares, err := slicepares(s)

	if err != nil {
		fmt.Print("Erro:", err)
	} else {
		fmt.Println("Números pares:", pares)
	}
}

func slicepares(s []int) ([]int, error) {

	if len(s) == 0 {
		return nil, fmt.Errorf("O slice está vazio")
	}
	var pares []int
	for _, num := range s {
		if num%2 == 0 {
			pares = append(pares, num)
		}
	}
	return pares, nil
}
