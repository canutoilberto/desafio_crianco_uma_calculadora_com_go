package main

import (
	"fmt"
)

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}

func divide(a, b int) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return float64(a) / float64(b), nil
}

func main() {
	var operation string
	var a, b int

	fmt.Print("Enter an operation (+, -, *, /): ")
	fmt.Scanln(&operation)

	fmt.Print("Enter the first number: ")
	fmt.Scanln(&a)

	fmt.Print("Enter the second number: ")
	fmt.Scanln(&b)

	var result interface{}
	var err error

	switch operation {
	case "+":
		result = add(a, b)
	case "-":
		result = subtract(a, b)
	case "*":
		result = multiply(a, b)
	case "/":
		result, err = divide(a, b)
		if err != nil {
			fmt.Println(err)
			return
		}
	default:
		fmt.Println("Invalid operation")
		return
	}

	fmt.Printf("Result: %v\n", result)
}
