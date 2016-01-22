package main

import "fmt"

func main() {
    pos, result := 0, 0
    for i := 0; i <= 10; i++ {
        pos, result = fibonacci(i)
        fmt.Printf("[%d] %d\n", pos, result)
    }
}

func fibonacci(n int) (pos, res int) {
    if n <= 1 {
        res = 1
    } else {
        _, res1 := fibonacci(n-2)
        _, res2 := fibonacci(n-1)
        res = res1 + res2
    }
    pos = n
    return
}
