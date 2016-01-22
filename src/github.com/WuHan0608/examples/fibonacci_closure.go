package main

import "fmt"

func main() {
    for i := 0; i <= 10; i++ {
        fmt.Printf("%d ", Fib(i))
    }
    fmt.Println()
}

func Fib() func(int) int {
    var x int
    return func(b int) int {
        if a <= 1 {
            x = 1 
        } else {
            x += b
        }
    }
}
