package main

import "fmt"

func main() {
    var i int = 5
    for i > 0 {
        i--
        fmt.Printf("The variable is now: %d\n", i)
    }
}
