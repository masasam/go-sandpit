package main

import "fmt"

func main() {
	x := 1
	{
		y := 2
		{
			z := 3
			fmt.Println(x)
			fmt.Println(y)
			fmt.Println(z)
		}
		fmt.Println(x)
		fmt.Println(y)
		// fmt.Println(z) error
	}
	fmt.Println(x)
	//fmt.Println(y) error
	//fmt.Println(z) error
}
