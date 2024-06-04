package main

import "fmt"

/*
Дана структура Human (с произвольным набором полей и методов). Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
*/

type Human struct {
	Name string
	Age  int
	Sex  rune
}

func (h *Human) Walk() {
	fmt.Println(h.Name, "Walking")
}
func (h *Human) Eat() {
	fmt.Println(h.Name, "Eating")
}
func (h *Human) Sing() {
	fmt.Println(h.Name, "Singing")
}

type Action struct {
	Human Human
}

func main() {
	action := Action{
		Human: Human{
			Name: "Zeleboba",
			Age:  17,
			Sex:  'F',
		},
	}

	action.Human.Eat()
	action.Human.Walk()
	action.Human.Walk()
}
