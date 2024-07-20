package functions

import "fmt"

// Function that returns a greeting message
func greet(name string) string {
	return "Hello, " + name
}

// Function that adds two integers
func add(a int, b int) int {
	return a + b
}

// Function that returns the square of an integer
func square(x int) int {
	return x * x
}

// Function with multiple parameters and return values
func divide(dividend, divisor int) (int, int) {
	quotient := dividend / divisor
	remainder := dividend % divisor
	return quotient, remainder
}

// Function with named return values
func multiply(x, y int) (result int) {
	result = x * y
	return // Named return value
}

// Variadic function that sums any number of integers
func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

// Function literal (anonymous function) assigned to a variable
func Functions() {
	// Define and call a function literal
	square := func(x int) int {
		return x * x
	}
	fmt.Println("Function Literals:")
	fmt.Println(square(4))

	// Display function definitions and usages
	fmt.Println("Defining Functions:")
	fmt.Println(greet("Alice"))

	fmt.Println("Function Parameters:")
	fmt.Println(add(5, 3))

	fmt.Println("Return Values:")
	fmt.Println(square(4))

	// Demonstrate multiple return values
	quotient, remainder := divide(10, 3)
	fmt.Printf("Divide Function: Quotient = %d, Remainder = %d\n", quotient, remainder)

	// Demonstrate named return values
	fmt.Println("Multiply Function (Named Return Value):")
	fmt.Println(multiply(6, 7))

	// Demonstrate variadic function
	fmt.Println("Sum Function (Variadic):")
	fmt.Println(sum(1, 2, 3, 4, 5))

	// Function literal demonstration
	fmt.Println("Function Literal:")
	fmt.Println(square(10))
}
