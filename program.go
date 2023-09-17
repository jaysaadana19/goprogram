package main

import (
    "fmt"
)

// Function to calculate the factorial of a number
func factorial(n int) uint64 {
    if n == 0 || n == 1 {
        return 1
    }
    fact := uint64(1)
    for i := 2; i <= n; i++ {
        fact *= uint64(i)
    }
    return fact
}

func main() {
    var num int

    // Input
    fmt.Print("Enter a non-negative integer: ")
    fmt.Scanf("%d", &num)

    if num < 0 {
        fmt.Println("Factorial is undefined for negative numbers.")
    } else {
        // Calculate and display the factorial
        fact := factorial(num)
        fmt.Printf("Factorial of %d is %d\n", num, fact)
    }
}
