package main

import "fmt"

func findIntersection[T comparable](arr1, arr2 []T) []T {
	intersaction := []T{}

	m := make(map[T]struct{})

	for _, value := range arr1 {
		m[value] = struct{}{}
	}

	for _, value := range arr2 {
		if _, exists := m[value]; exists {
			intersaction = append(intersaction, value)
		}
	}
	return intersaction
}

func test1() {
	arr1 := []string{"cat", "dot", "pog", "ping"}
	arr2 := []string{"cat", "pong", "pog"}

	res := findIntersection(arr1, arr2)
	fmt.Println(res)
}
func test2() {
	arr1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	arr2 := []int{1, 3, 5, 7, 9, 11}

	res := findIntersection(arr1, arr2)
	fmt.Println(res)
}

func test3() {
	arr1 := []rune{'t', 'u', 'r', 'i', 'n', 'g', 't', 'e', 'a', 'm', 's', 'p', 'i', 'r', 'i', 't'}
	arr2 := []rune{'t', 'u', 'r', 'i', 'n', 'g'}

	res := findIntersection(arr1, arr2)
	fmt.Println(string(res))
}

func main() {
	test1()
	test2()
	test3()
}
