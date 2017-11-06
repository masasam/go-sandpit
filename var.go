package main

import "fmt"

var x = 10

func square(x int) int {
	return x * x
}

func main() {
	fmt.Println(x)
	fmt.Println(square(5))
	fmt.Println(x)
}
