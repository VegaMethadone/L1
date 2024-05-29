package main

import (
	"fmt"
	"strings"
)

func main() {
	var builder strings.Builder

	builder.WriteString("Hello")
	builder.WriteString(" ")
	builder.WriteString("world")
	builder.WriteString("!")

	result := builder.String()
	fmt.Println(result)
}
