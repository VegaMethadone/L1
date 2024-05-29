package main

import (
	"fmt"
	"sync"
)

type HashMap struct {
	mu sync.RWMutex
	m  map[int]int
}

func newHashMap() HashMap {
	return HashMap{m: make(map[int]int)}
}

func (h *HashMap) add(x int) {
	h.mu.Lock()
	h.m[x]++
	h.mu.Unlock()
}

func (h *HashMap) find(x int) {
	h.mu.RLock()
	if value, exists := h.m[x]; exists {
		fmt.Printf("key: %d, value: %d\n", x, value)
	} else {
		fmt.Println(x, "not found")
	}
	h.mu.RUnlock()
}

func writeData(hash *HashMap, wg *sync.WaitGroup, amount int) {
	defer wg.Done()
	for i := 0; i < amount; i++ {
		hash.add(i)
	}
}

func main() {
	hash := newHashMap()

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go writeData(&hash, &wg, 1001)
	}
	wg.Wait()

	hash.find(1000)
	fmt.Println("Done!")
}
