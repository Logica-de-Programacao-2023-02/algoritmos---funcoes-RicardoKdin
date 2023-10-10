package main

import "fmt"

func main() {

	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{1, 3, 5, 7, 9}
	ambos, err := ambosslice(s1, s2)

	if err != nil {
		fmt.Print("Erro: ", err)
	} else {
		fmt.Println(ambos)
	}

}

func ambosslice(s1, s2 []int) ([]int, error) {

	if len(s1) == 0 || len(s2) == 0 {
		return nil, fmt.Errorf("Slice vazio")
	}

	mapaS1 := make(map[int]struct{})
	for _, num := range s1 {
		mapaS1[num] = struct{}{}
	}

	var ambos []int
	for _, num := range s2 {
		if _, existe := mapaS1[num]; existe {
			ambos = append(ambos, num)
		}
	}
	return ambos, nil
}
