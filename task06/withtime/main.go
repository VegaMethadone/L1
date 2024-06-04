package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var globalSeed = rand.New(rand.NewSource(time.Now().Unix()))

func worker(id int, ctx context.Context, out <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d is stoprd\n", id)
			return
		case value := <-out:
			fmt.Printf("Worker %d got: %d\n", id, value)
		}
	}
}

func writeData(ctx context.Context, in chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Chan is closed")
			close(in)
			return
		default:
			value := globalSeed.Int() % 1000
			in <- value
			fmt.Println("Wtire data:", value)
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func main() {
	ch := make(chan int, 10)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(4))
	var wg sync.WaitGroup

	wg.Add(1)
	go writeData(ctx, ch, &wg)

	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go worker(i, ctx, ch, &wg)
	}

	time.Sleep(time.Second * 6)
	cancel()
	wg.Wait()

	fmt.Println("Done")
}
