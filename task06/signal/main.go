package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var globalSeed = rand.New(rand.NewSource(time.Now().Unix()))

func worker(id int, ctx context.Context, out <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d is stopped\n", id)
			return
		case value := <-out:
			fmt.Printf("Worker %d got data: %d\n", id, value)
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
			fmt.Println("Add value:", value)
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func main() {
	ch := make(chan int, 10)
	stop := make(chan os.Signal, 1)

	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	wg.Add(1)
	go writeData(ctx, ch, &wg)

	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go worker(i, ctx, ch, &wg)
	}

	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	cancel()

	wg.Wait()
	fmt.Println("Done")
}
