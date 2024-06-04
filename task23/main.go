package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	newArr := deleteElemFromSlice(arr, 5)
	fmt.Println(newArr)
}

func deleteElemFromSlice(arr []int, elem int) []int {
	newArr := make([]int, len(arr)-1)
	if elem == 1 {
		copy(newArr, arr[1:])
	} else if elem == len(arr) {
		copy(newArr, arr[:len(arr)-1])
	} else {
		copy(newArr, append(arr[:elem-1], arr[elem:]...))
	}
	return newArr
}
