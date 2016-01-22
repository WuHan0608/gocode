package main

import "fmt"

func main() {
<<<<<<< HEAD
    for i := 0; i <= 10; i++ {
        fmt.Printf("%d ", Fib(i))
=======
    f := fib()
    for i := 0; i <= 10; i++ {
        fmt.Printf("%d ", f())
>>>>>>> e8eba3a542992d0a6b17788183e5e2d32bebea8a
    }
    fmt.Println()
}

<<<<<<< HEAD
func Fib() func(int) int {
    var x int
    return func(b int) int {
        if a <= 1 {
            x = 1 
        } else {
            x += b
        }
=======
func fib() func() int {
    var x, y int = 0, 0
    return func() int {
        if y == 0 {
            y = 1
        } else {
            x, y = y, x+y
        }
        return y
>>>>>>> e8eba3a542992d0a6b17788183e5e2d32bebea8a
    }
}
