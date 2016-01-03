package main

import (
    "fmt"
    "os"
)

func main() {
    var goos string = os.Getenv("GOPATH")
    fmt.Printf("GOPATH is: %s\n", goos)
    path := os.Getenv("PATH")
    fmt.Printf("Path is: %s\n", path)
}
