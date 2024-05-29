package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»). Символы могут быть unicode.
*/

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println()
		line, _ := reader.ReadString('\n')
		reverseWord(line)
	}

}

func reverseWord(str string) {
	newRune := []rune(str)
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		newRune[i], newRune[j] = newRune[j], newRune[i]
	}
	fmt.Println(string(newRune))
}
