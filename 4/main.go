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

/*
Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
Необходима возможность выбора количества воркеров при старте.

Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.
*/

var globalSeed = rand.New(rand.NewSource(time.Now().Unix()))

func randomeNumber() int {
	return globalSeed.Int() % 1000
}

func worker(ctx context.Context, id int, out <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d is stoped\n", id)
			return
		case value := <-out:
			fmt.Printf("Worker %d received value: %d\n", id, value)
		}
	}
}

func sendData(ctx context.Context, in chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			close(in)
			fmt.Println("Main chan is closed")
			return
		case in <- randomeNumber():
			time.Sleep(time.Millisecond * 500)
		}
	}

}

func main() {
	var n int
	fmt.Println("Input N number of workers:")
	fmt.Scan(&n)

	var wg sync.WaitGroup
	ch := make(chan int, 10)
	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)
	go sendData(ctx, ch, &wg)

	for id := 1; id <= n; id++ {
		wg.Add(1)
		go worker(ctx, id, ch, &wg)
	}

	sign := make(chan os.Signal, 1)
	signal.Notify(sign, os.Interrupt, syscall.SIGTERM)

	<-sign
	cancel()

	wg.Wait()

	fmt.Println("All workers are stoped")
}
