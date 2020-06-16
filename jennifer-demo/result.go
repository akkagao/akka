package main

import "fmt"

func main() {
	var a = 10
	var b = 20
	fmt.Println(demo(a, b))
	fmt.Println("Hello, world")
}

func demo(a, b int) int {
	return a + b
}
