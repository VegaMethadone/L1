package main

import (
	"fmt"
	"slices"
	"sort"
)

/*
Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
*/

func main() {
	arr := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	sort1(arr)
	sort2(arr)
	sort3(arr)
}

func sort1(arr []int) {
	newArr := make([]int, len(arr))
	copy(newArr, arr)
	slices.Sort(newArr)
	fmt.Println(newArr)
}

func sort2(arr []int) {
	newArr := make([]int, len(arr))
	copy(newArr, arr)
	sort.Ints(newArr)
	fmt.Println(newArr)
}

func sort3(arr []int) {
	newArr := make([]int, len(arr))
	copy(newArr, arr)
	sort.Slice(newArr, func(i, j int) bool {
		return newArr[i] < newArr[j]
	})
	fmt.Println(newArr)
}

func quikeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[0]
	left := []int{}
	right := []int{}

	for _, value := range arr[1:] {
		if value < pivot {
			left = append(left, value)
		} else {
			right = append(right, value)
		}
	}
	left = quikeSort(left)
	right = quikeSort(right)

	return append(append(left, pivot), right...)
}
