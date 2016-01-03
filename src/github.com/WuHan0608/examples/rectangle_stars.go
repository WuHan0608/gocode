package main

import "fmt"

const (
    WIDTH = 20
    HEIGHT = 10
)

func main() {
    for i := 0; i < HEIGHT; i++ {
        fmt.Printf("*")
        for j := 0; j < WIDTH; j++ {
            switch {
                case i == 0 || i == HEIGHT-1:
                    fmt.Printf("*")
                default:
                    fmt.Printf(" ")
            }
        }
        fmt.Println("*")
    }
}
