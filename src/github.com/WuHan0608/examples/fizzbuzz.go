package main

import "fmt"

const (
    FIZZ = 3
    BUZZ = 5
    FIZZBUZZ = 15
)

func main() {
    for i := 1; i <= 100; i++ {
        switch {
            case i % FIZZBUZZ == 0:
                fmt.Println("FizzBuzz")
            case i % BUZZ == 0:
                fmt.Println("Buzz")
            case i % FIZZ == 0:
                fmt.Println("Fizz")
            default:
                fmt.Println(i)
        }
    }
}
