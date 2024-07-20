package control_structures

import "fmt"

// Function demonstrating various ways to use if-else statements
func IfElse() {
	fmt.Println("If-Else:")

	age := 18
	hasPermission := true

	// 1. Basic if-else statement
	if age >= 18 && hasPermission {
		fmt.Println("Adult with permission")
	} else if age >= 18 {
		fmt.Println("Adult without permission")
	} else {
		fmt.Println("Minor")
	}

	// 2. Simple if statement with a block
	if age >= 18 {
		fmt.Println("Person is an adult.")
	}

	// 3. If-else with an initialization statement
	if discount := 10; discount > 0 {
		fmt.Println("Discount available:", discount)
	} else {
		fmt.Println("No discount available")
	}

	// 4. Nested if statements
	if age >= 18 {
		if hasPermission {
			fmt.Println("Adult with permission")
		} else {
			fmt.Println("Adult without permission")
		}
	} else {
		fmt.Println("Minor")
	}

	// 5. If statement with a condition directly
	if age >= 18 {
		fmt.Println("Adult")
	}

	// 6. Multiple conditions with logical operators
	if age >= 18 && hasPermission {
		fmt.Println("Adult with permission")
	} else if age >= 18 || hasPermission {
		fmt.Println("Adult or has permission")
	} else {
		fmt.Println("Neither adult nor has permission")
	}

	// 7. If with multiple conditions and block
	if age >= 18 && hasPermission {
		fmt.Println("Adult with permission")
	} else if age < 18 {
		fmt.Println("Minor")
	} else {
		fmt.Println("Adult without permission")
	}

	// 8. Switch-like behavior with if-else if blocks
	if age < 13 {
		fmt.Println("Child")
	} else if age < 18 {
		fmt.Println("Teen")
	} else if age < 60 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Senior")
	}
}
