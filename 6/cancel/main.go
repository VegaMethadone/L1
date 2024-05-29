package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var globalSeed = rand.New(rand.NewSource(time.Now().Unix()))

func writeData(ctx context.Context, in chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Ch is closed")
			close(in)
			return
		default:
			value := rand.Intn(1000)
			in <- value
			fmt.Println("added:", value)
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func worker(id int, ctx context.Context, out <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d is stoped\n", id)
			return
		case value := <-out:
			fmt.Printf("Worker %d got: %d\n", id, value)
		}
	}

}

func main() {
	ch := make(chan int, 10)

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)
	go writeData(ctx, ch, &wg)

	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go worker(i, ctx, ch, &wg)
	}

	time.Sleep(time.Second * 4)
	cancel()

	wg.Wait()
	fmt.Println("Progrm is done")
}
