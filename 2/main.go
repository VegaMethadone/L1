package main

import (
	"fmt"
	"sync"
)

/*
Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/

func main() {
	arr := []int{2, 4, 6, 8, 10}
	mid := len(arr) / 2

	arr1 := arr[:mid]
	arr2 := arr[mid:]

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i, value := range arr1 {
			arr1[i] = value * value
		}
	}()

	go func() {
		defer wg.Done()
		for i, value := range arr2 {
			arr2[i] = value * value
		}
	}()
	wg.Wait()
	fmt.Println(arr)
}
