package main

import "fmt"

/*
Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
*/

func main() {
	var num int64 = 64
	fmt.Println(chancheBit(num, true, 1))  // 65
	fmt.Println(chancheBit(num, false, 7)) //  0
}

func chancheBit(num int64, isOne bool, pos uint) int64 {
	pos -= 1
	if isOne {
		num |= 1 << pos
	} else {
		num &^= 1 << pos
	}
	return num
}
