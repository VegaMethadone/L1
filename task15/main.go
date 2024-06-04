package main

import "fmt"

func createHugeString(n int) []byte {
	str := make([]byte, n)
	for i := range str {
		str[i] = '_'
	}
	return str
}

func someFunc() {
	v := createHugeString(1 << 10)
	justString := v[:100]
	str := make([]byte, 100)
	copy(str, justString)
	str[0], str[99] = '0', '0'

	fmt.Println("Original: ", string(justString))
	fmt.Println("New: ", string(str))
}

func main() {
	someFunc()
}
