package main

import "fmt"

/*
Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

Например:
abcd — true
abCdefAaf — false
	aabcd — false
*/

func main() {
	var str string

	for {
		fmt.Scan(&str)
		fmt.Println(checkUnique(str))
	}
}

func checkUnique(str string) bool {
	m := make(map[rune]bool)
	for _, char := range str {
		if m[char] {
			return false
		}
		m[char] = true
	}
	return true
}
