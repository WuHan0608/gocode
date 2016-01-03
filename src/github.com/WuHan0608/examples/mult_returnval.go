package main

import "fmt"

func main() {
    var sum, product, delta int
    num1, num2 := 10, 20

    // unnamed return variables
    sum, product, delta = sum_product_delta_1(num1, num2)
    fmt.Printf("%d + %d = %d, %d * %d = %d, %d - %d = %d\n", num1, num2, sum, num1, num2, product, num1, num2, delta)

    // named return variables
    sum, product, delta = sum_product_delta_2(num1, num2)
    fmt.Printf("%d + %d = %d, %d * %d = %d, %d - %d = %d\n", num1, num2, sum, num1, num2, product, num1, num2, delta)

}

func sum_product_delta_1(n1, n2 int) (int, int, int) {
    return n1 + n2, n1 * n2, n1 - n2
}

func sum_product_delta_2(n1, n2 int) (sum, product, delta int) {
    sum = n1 + n2
    product = n1 * n2
    delta = n1 - n2
    return
}
