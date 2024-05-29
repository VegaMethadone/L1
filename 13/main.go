package main

import "fmt"

func swap1(x, y *int) {
	*x = *x + *y
	*y = *x - *y
	*x = *x - *y

}

func swap2(x, y *int) {
	*x = *x ^ *y
	*y = *x ^ *y
	*x = *x ^ *y
}

func main() {
	x := 10
	y := 5
	fmt.Printf("x: %d, y: %d\n", x, y)
	swap2(&x, &y)
	fmt.Printf("x: %d, y: %d\n", x, y)
}
