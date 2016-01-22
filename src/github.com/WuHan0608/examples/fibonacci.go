package main

<<<<<<< HEAD
import "fmt"

func main() {
    result := 0
    for i := 0; i <= 10; i++ {
        result = fibonacci(i)
        fmt.Printf("fibonacci(%d) is: %d\n", i, result)
    }
=======
import (
    "fmt"
    "time"
)

func main() {
    result := 0
    start := time.Now()
    for i := 0; i <= 25; i++ {
        result = fibonacci(i)
        fmt.Printf("fibonacci(%d) is: %d\n", i, result)
    }
    end := time.Now()
    delta := end.Sub(start)
    fmt.Printf("longCalculation took this amount of time: %s\n", delta)
>>>>>>> e8eba3a542992d0a6b17788183e5e2d32bebea8a
}

func fibonacci(n int) (res int) {
    if n <= 1 {
        res = 1
    } else {
        res = fibonacci(n-1) + fibonacci(n-2)
    }
    return
}
