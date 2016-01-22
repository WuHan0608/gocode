package main

import (
    "fmt"
    "strings"
)

func main() {
    assiiOnly := func(c rune) rune { if c > 127 { return '?' }; return c }
    fmt.Println(strings.Map(assiiOnly, "Jérôme Österreich"))
}
