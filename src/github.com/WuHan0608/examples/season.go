package main

import "fmt"

func main() {
    for i := 1; i < 13; i++ {
        fmt.Printf("%d: ", i)
        switch {
            case i < 4:
                fmt.Println("Spring")
            case i < 7:
                fmt.Println("Summer")
            case i < 10:
                fmt.Println("Autumn")
            default:
                fmt.Println("Winter")
        }
    }
}
