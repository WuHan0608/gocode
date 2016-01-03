package main

import "fmt"

func main() {
//For-Loop
    fmt.Println("For-Loop:")
    for i := 0; i < 15; i++ {
        fmt.Printf(" counter: %d\n", i)
    }

// Goto Loop
    fmt.Println("Goto Loop:")
    i := 0
LOOP:
    fmt.Printf(" counter: %d\n", i)
    i++
    if i < 15 {
        goto LOOP
    }
}
