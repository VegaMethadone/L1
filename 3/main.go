package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.
*/

func main() {
	arr := []int{2, 4, 6, 8, 10}
	mid := len(arr) / 2
	var wg sync.WaitGroup
	var sum atomic.Int64

	arr1 := arr[:mid]
	arr2 := arr[mid:]

	wg.Add(2)

	go func() {
		defer wg.Done()
		for _, value := range arr1 {
			sum.Add(int64(value * value))
		}
	}()
	go func() {
		defer wg.Done()
		for _, value := range arr2 {
			sum.Add(int64(value * value))
		}
	}()

	wg.Wait()
	fmt.Println(sum.Load())
}
