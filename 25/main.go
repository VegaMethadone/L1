package main

import (
	"context"
	"fmt"
	"time"
)

func sleep1(x int) {
	start := time.Now()
	<-time.After(time.Second * time.Duration(x))
	fmt.Println("Sleep: ", time.Now().Sub(start))
}

func sleep2(x int) {
	start := time.Now()
	for {
		if time.Since(start) >= time.Duration(x)*time.Second {
			break
		}
	}
	fmt.Println("Sleep: ", time.Now().Sub(start))
}

func sleep3(x int) {
	start := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(x)*time.Second)
	defer cancel()

	select {
	case <-ctx.Done():
		fmt.Println("Sleep: ", time.Now().Sub(start))
		return
	}
}

func main() {
	var n, x int
	fmt.Println("How many iterations?")
	fmt.Scan(&n)
	fmt.Println("How long to sleep?")
	fmt.Scan(&x)

	for i := 0; i < n; i++ {
		fmt.Println("test1:", i)
		sleep1(x)
	}
	for i := 0; i < n; i++ {
		fmt.Println("test2:", i)
		sleep2(x)
	}
	for i := 0; i < n; i++ {
		fmt.Println("test3 :", i)
		sleep3(x)
	}
}
