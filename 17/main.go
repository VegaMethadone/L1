package main

import (
	"fmt"
	"sort"
)

var arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func main() {
	search2()
}

func search1() {
	fmt.Println("Put number [1:10]:")
	var x int
	for {
		fmt.Scan(&x)
		index := sort.Search(len(arr), func(i int) bool { return arr[i] >= x })
		if index < len(arr) && arr[index] == x {
			fmt.Printf("Found %d at pos %d\n", x, index)
		} else {
			fmt.Println("Not found")
		}
	}
}

func search2() {
	fmt.Println("Put number [1:10]:")
	var x int
	for {
		fmt.Scan(&x)
		index := sort.SearchInts(arr, x)
		if index < len(arr) && arr[index] == x {
			fmt.Printf("Found %d at pos %d\n", x, index)
		} else {
			fmt.Println("Not found")
		}
	}
}
