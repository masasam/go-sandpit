package main

import "fmt"

func even(n int) bool {
	if n%2 == 0 {
		return true
	}
	return false
}

func removeif(f func(int) bool, ary []int) []int {
	buff := make([]int, 0)
	for _, v := range ary {
		if !f(v) {
			buff = append(buff, v)
		}
	}
	return buff
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(removeif(even, a))
}
