package main

import "fmt"

func insertSort(ary []int) {
	for i := 1; i < len(ary); i++ {
		tmp := ary[i]
		j := i - 1
		for ; j >= 0 && tmp < ary[j]; j-- {
			ary[j+1] = ary[j]
		}
		ary[j+1] = tmp
	}
}

func main() {
	a := []int{5, 6}
	b := []int{9, 1, 7, 2, 9, 8, 3, 2, 1, 0}
	c := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	insertSort(a)
	insertSort(b)
	insertSort(c)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
