package main

import "fmt"

func main() {
    recursive_print(1)
    fmt.Println()
}

func recursive_print(n int) {
    if n > 10 {
        return
    }
    recursive_print(n+1)
    fmt.Printf("%d ", n)
}
