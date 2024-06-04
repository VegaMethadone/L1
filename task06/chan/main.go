package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var globalSeed = rand.New(rand.NewSource(time.Now().Unix()))

func worker(id int, stop <-chan bool, out <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-stop:
			fmt.Printf("Worker %d is finished\n", id)
			return
		case value := <-out:
			fmt.Printf("Worker %d got:  %d\n", id, value)
		}
	}
}

func writeData(stop <-chan bool, in chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-stop:
			fmt.Println("chan is closed")
			close(in)
			return
		default:
			value := globalSeed.Int() % 1000
			in <- value
			fmt.Printf("Chan got value: %d\n", value)
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func main() {
	ch := make(chan int, 2)
	stop := make(chan bool, 1)

	var wg sync.WaitGroup

	wg.Add(1)
	go writeData(stop, ch, &wg)

	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go worker(i, stop, ch, &wg)
	}

	time.Sleep(time.Second * 4)
	close(stop)

	wg.Wait()
	fmt.Println("Done")
}
