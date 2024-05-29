package main

import "fmt"

type Pair struct {
	a int64
	b int64
}

func NewPair(a, b int64) Pair {
	return Pair{a: a, b: b}
}

func (p *Pair) Add() int64 {
	return p.a + p.b
}

func (p *Pair) Subtract() int64 {
	return p.a - p.b
}

func (p *Pair) Divide() int64 {
	return p.a / p.b
}

func (p *Pair) Multiply() int64 {
	return p.a * p.b
}

func main() {
	test := NewPair(int64(1<<30+1), int64(1<<26+1))

	fmt.Println("Add", test.Add())
	fmt.Println("Subtract", test.Subtract())
	fmt.Println("Divide", test.Divide())
	fmt.Println("Multiply", test.Multiply())
}
