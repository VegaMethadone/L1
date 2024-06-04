package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 1
	b := "pog"
	c := true
	d := make(chan int)

	typeDefinitor(a)
	typeDefinitor(b)
	typeDefinitor(c)
	typeDefinitor(d)
}

func typeDefinitor(x interface{}) {
	res := reflect.TypeOf(x)
	switch res.Kind() {
	case reflect.Int:
		fmt.Printf("Value of {%v} is %v\n", x, reflect.Int)
	case reflect.Bool:
		fmt.Printf("Value of {%v} is %v\n", x, reflect.Bool)
	case reflect.String:
		fmt.Printf("Value of {%v} is %v\n", x, reflect.String)
	case reflect.Chan:
		fmt.Printf("Value of {%v} is %v\n", x, reflect.Chan)
	default:
		fmt.Println("Another type")
	}
}
