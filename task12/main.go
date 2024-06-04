package main

import "fmt"

/*
Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
*/

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	newArr := []string{}
	createOwnUnity(&arr, &newArr)

	fmt.Println(newArr)
}

func createOwnUnity(arr, newArr *[]string) {
	m := make(map[string]bool)
	for _, value := range *arr {
		if m[value] {
			continue
		} else {
			(*newArr) = append((*newArr), value)
		}
		m[value] = true
	}
}
