package main

import "fmt"

func main() {
    f := fib()
    for i := 0; i <= 10; i++ {
        fmt.Printf("%d ", f())
    }
    fmt.Println()
}

func fib() func() int {
    var x, y int = 0, 0
    return func() int {
        if y == 0 {
            y = 1
        } else {
            x, y = y, x+y
        }
        return y
    }
}
