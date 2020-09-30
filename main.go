package main

import "fmt"

func main() {
	keys := []string{}

	carta := map[string][]int{
		"1": {1, 2},
		"2": {3, 4},
	}

	//Перебор мапа
	for _, val := range carta {
		// keys = append(keys, key)
		fmt.Println(val)
	}
	fmt.Println(keys)

	//Создание мапа через мейк
	h := make(map[int]string)
	fmt.Println(h)
}
