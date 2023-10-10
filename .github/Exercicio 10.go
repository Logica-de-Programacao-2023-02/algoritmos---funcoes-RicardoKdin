package main

import (
	"fmt"
	"sort"
)

func main() {

	s := []int{1, 3, 5, 7, 9, 8, 6, 4, 2}
	cres, err := crescente(s)

	if err != nil {
		fmt.Print("Erro: ", err)
	} else {
		fmt.Println("Números crescentes: ", cres)
	}

}

func crescente(s []int) ([]int, error) {

	if len(s) == 0 {
		return nil, fmt.Errorf("A string está vazia")
	}
	novoslice := make([]int, len(s))
	copy(novoslice, s)

	sort.Ints(novoslice)

	return novoslice, nil
}
