package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var globalRand = rand.New(rand.NewSource(time.Now().Unix()))

func randomNumber() int {
	return globalRand.Int() % 100
}

func createArr(size int) []int {
	arr := make([]int, size)
	for i := range arr {
		arr[i] = randomNumber()
	}
	return arr
}

func writeData(in chan<- int, wg *sync.WaitGroup, arr *[]int) {
	defer wg.Done()

	for _, value := range *arr {
		fmt.Println("Chan 1", value)
		in <- value
	}

	close(in)
}

func transitData(in chan<- int, out <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for value := range out {
		fmt.Println("chan 2:", value*value)
		in <- value * value
	}

	close(in)
}

func readData(out <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case value, ok := <-out:
			if !ok {
				fmt.Println("Chan 2 is closed")
				return
			}
			fmt.Println("Got:", value)
		}
	}
}

func main() {
	var n int
	fmt.Println("Put the size of arr:")
	fmt.Scan(&n)

	arr := createArr(n)
	var wg sync.WaitGroup

	ch1 := make(chan int, 4)
	ch2 := make(chan int, 4)

	wg.Add(3)
	go writeData(ch1, &wg, &arr)
	go transitData(ch2, ch1, &wg)
	go readData(ch2, &wg)
	wg.Wait()

	fmt.Println("Done")

}
