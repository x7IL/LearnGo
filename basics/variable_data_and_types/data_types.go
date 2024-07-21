package variable_data_and_types

import "fmt"

// Global variable
var globalVar int = 100

// Define a structure
type Person struct {
	Name string
	Age  int
}

// Function demonstrating various ways to declare variables
func VariableExamples() {
	// 1. Declare and initialize with `var`
	var x int = 10
	var y = 20 // Type inferred as int

	// 2. Declare multiple variables
	var a, b int = 1, 2

	// 3. Declare with zero value
	var z int // z is initialized to 0

	// 4. Using short variable declaration
	localVar := 30
	names := []string{"Alice", "Bob"}

	// 5. Declare constants
	const Pi = 3.14
	const (
		A = 1
		B = 2
	)

	// 6. Declare with struct
	p := Person{Name: "Alice", Age: 30}

	// 7. Declare with pointers and maps
	var ptr *int         // Pointer, initialized to nil
	var m map[string]int // Map, initialized to nil

	// Print values
	fmt.Println("Variables:")
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("z:", z)
	fmt.Println("localVar:", localVar)
	fmt.Println("names:", names)
	fmt.Println("Pi:", Pi)
	fmt.Println("A:", A)
	fmt.Println("B:", B)
	fmt.Println("Person:", p)
	fmt.Println("Pointer:", ptr)
	fmt.Println("Map:", m)
}
