package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func addData(ctx context.Context, ch chan<- int, m int) {
	defer close(ch)
	for i := 0; i < m; i++ {
		select {
		case <-ctx.Done():
			return
		case ch <- i:
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func readData(ctx context.Context, ch <-chan int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Timeout")
			return
		case value := <-ch:
			fmt.Println("Reading:", value)
		}
	}
}

func main() {
	var m, n int
	fmt.Println("Input amount(M) of operations")
	fmt.Scan(&m)
	fmt.Println("Input time(N) for operations")
	fmt.Scan(&n)

	ch := make(chan int, 10)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(n)*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		addData(ctx, ch, m)
	}()
	go func() {
		defer wg.Done()
		readData(ctx, ch)
	}()

	wg.Wait()
	fmt.Println("Done")
}
