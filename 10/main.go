package main

import "fmt"

/*
Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
Объединить данные значения в группы с шагом в 10 градусов.
Последовательность в подмножноствах не важна.
*/

func main() {
	m := make(map[int][]float64)

	var n float64
	for {
		fmt.Scan(&n)
		result := n / 10
		m[int(result)*10] = append(m[int(result)*10], n)

		for key, value := range m {
			fmt.Printf("UNION: %d, values: %v\n", key, value)
		}
	}
}
