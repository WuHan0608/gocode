package main

import (
    "fmt"
)

func main() {
    n := 30
    res := factorial(30)
    fmt.Printf("%d! = %d\n", n, res)
}

func factorial(n int) (res uint64) {
    if n == 0 {
        res = 1
    } else {
        res = uint64(n) * factorial(n-1)
    }
    return
}
