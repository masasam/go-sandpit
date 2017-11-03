package main

import (
    "fmt"
    "time"
)

func tarai(x, y, z int) int {
    if x <= y {
        return z
    } else {
        return tarai(tarai(x - 1, y, z), tarai(y - 1, z, x), tarai(z - 1, x, y))
    }
}

func main() {
    s := time.Now()
    fmt.Println(tarai(22, 11, 0))
    e := time.Now().Sub(s)
    fmt.Println(e)
}
