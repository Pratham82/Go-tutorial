package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// add takes two float64 numbers and returns their sum
func add(a, b float64) float64 {
	return a + b
}

// subtract takes two float64 numbers and returns their difference
func subtract(a, b float64) float64 {
	return a - b
}

// multiply takes two float64 numbers and returns their product
func multiply(a, b float64) float64 {
	return a * b
}

// divide takes two float64 numbers and returns their quotient
// Note: division by zero is handled in the main function
func divide(a, b float64) float64 {
	return a / b
}

// modulus takes two float64 numbers and returns the remainder
func modulus(a, b float64) float64 {
	return float64(int(a) % int(b))
}

func main() {
	// Create a reader to read user input from console
	// bufio.Reader is more efficient than fmt.Scan for multiple inputs
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("=== Go CLI Calculator ===")
	fmt.Println("Supported operations: +, -, *, /, %")
	fmt.Println("-------------------------")

	// Get first number from user
	fmt.Print("Enter first number: ")
	num1Str, _ := reader.ReadString('\n')
	// Remove whitespace/newline from input
	num1Str = strings.TrimSpace(num1Str)

	// Convert string to float64
	// strconv.ParseFloat returns two values: the float and an error
	num1, err := strconv.ParseFloat(num1Str, 64)
	if err != nil {
		fmt.Println("Error: Please enter a valid number")
		os.Exit(1) // Exit with error code 1
	}

	// Get operator from user
	fmt.Print("Enter operator (+, -, *, /, %): ")
	operator, _ := reader.ReadString('\n')
	operator = strings.TrimSpace(operator)

	// Validate operator
	// Only allow single character operators from our supported set
	if len(operator) != 1 {
		fmt.Println("Error: Please enter a valid operator")
		os.Exit(1)
	}

	// Get second number from user
	fmt.Print("Enter second number: ")
	num2Str, _ := reader.ReadString('\n')
	num2Str = strings.TrimSpace(num2Str)

	num2, err := strconv.ParseFloat(num2Str, 64)
	if err != nil {
		fmt.Println("Error: Please enter a valid number")
		os.Exit(1)
	}

	// Perform calculation based on operator
	// switch is cleaner than multiple if-else statements
	var result float64

	switch operator {
	case "+":
		result = add(num1, num2)
	case "-":
		result = subtract(num1, num2)
	case "*":
		result = multiply(num1, num2)
	case "/":
		// Check for division by zero
		if num2 == 0 {
			fmt.Println("Error: Cannot divide by zero")
			os.Exit(1)
		}
		result = divide(num1, num2)
	case "%":
		// Check for modulus by zero
		if num2 == 0 {
			fmt.Println("Error: Cannot calculate modulus with zero")
			os.Exit(1)
		}
		result = modulus(num1, num2)
	default:
		fmt.Println("Error: Invalid operator. Use +, -, *, /, or %")
		os.Exit(1)
	}

	// Print result with proper formatting
	// %.2f formats the number to 2 decimal places
	fmt.Printf("Result: %.2f\n", result)
}
