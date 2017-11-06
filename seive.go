package main

import (
	"fmt"
	"time"
  )

func getprime() {
    var primes [100]bool
    for n := 2; n < len(primes); n++ {
        if primes[n] { continue }
        for m := 2 * n; m < len(primes); m += n {
            primes[m] = true
        }
        fmt.Print(n, " ")
    }
}

func main() {
	s := time.Now()
	getprime()
	e := time.Now().Sub(s)
	fmt.Println(e)
}
