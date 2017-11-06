package main

import "fmt"

func tailfact(n, a int) int {
	if n == 0 {
		return a
	} else {
		return tailfact(n-1, a*n)
	}
}

func main() {
	for i := 0; i < 13; i++ {
		fmt.Println(i, ":", tailfact(i, 1))
	}
}
