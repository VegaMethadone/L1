package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Input line:")
		line, _ := reader.ReadString('\n')
		reverseWordsInLine(line)
	}
}

func reverseWordsInLine(str string) {
	str = strings.TrimSpace(str)
	arr := strings.Split(str, " ")

	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	fmt.Println(strings.Join(arr, " "))
}
