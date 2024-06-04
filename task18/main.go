package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде. По завершению программа должна выводить итоговое значение счетчика.
*/

func iterArr(arr []int, count *atomic.Int64, wg *sync.WaitGroup) {
	defer wg.Done()
	for range arr {
		count.Add(1)
	}
}

func main() {
	fmt.Println("Input size fo arr:")
	var size int
	fmt.Scan(&size)

	arr := make([]int, size)
	mid := len(arr) / 2
	var wg sync.WaitGroup
	var newCounter atomic.Int64

	wg.Add(2)

	go iterArr(arr[:mid], &newCounter, &wg)
	go iterArr(arr[mid:], &newCounter, &wg)

	wg.Wait()
	fmt.Println("Done!")
	fmt.Println("Counter: ", newCounter.Load())

}
