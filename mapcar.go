package main

import "fmt"

func plusone(n int) int {
	return n + 1
}

func mapcar(f func(int) int, ary []int) []int {
	buff := make([]int, len(ary))
	for i, v := range ary {
		buff[i] = f(v)
	}
	return buff
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(mapcar(plusone, a))
}
