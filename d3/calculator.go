package main

import "fmt"

func main() {
	var a, b float64
	var op string

	fmt.Print("Enter first number: ")
	fmt.Scan(&a)

	fmt.Print("Enter operator (+, -, *, /): ")
	fmt.Scan(&op)

	fmt.Print("Enter second number: ")
	fmt.Scan(&b)

	var result float64

	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			fmt.Println("Error: division by zero!")
			return
		}
		result = a / b
	default:
		fmt.Println("Unknown operator:", op)
		return
	}

	fmt.Printf("Result: %.2f\n", result)
}
