package main

import "fmt"

func main() {
	fmt.Println("✅ Go is working correctly!")

	a, b := 5, 3
	fmt.Printf("Sum of %d + %d = %d\n", a, b, a+b)

	for i := 1; i <= 3; i++ {
		fmt.Printf("Loop %d\n", i)
	}
}
