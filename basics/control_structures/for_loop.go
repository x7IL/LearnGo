package control_structures

import "fmt"

// Main function to demonstrate various for loop types
func ForLoop() {
	fmt.Println("For Loop:")

	// 1. Classic for loop with initialization, condition, and increment
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("For Loop - with single condition:")
	// 2. For loop with a single condition (without initialization or increment)
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	fmt.Println("For Loop - infinite loop with break:")
	// 3. Infinite for loop with a break statement
	i = 0
	for {
		if i >= 5 {
			break
		}
		fmt.Println(i)
		i++
	}

	fmt.Println("For Loop - using range with slices:")
	// 4. For loop using `range` to iterate over slices
	slice := []int{10, 20, 30, 40, 50}
	for index, value := range slice {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	fmt.Println("For Loop - using range with maps:")
	// 5. For loop using `range` to iterate over maps
	mymap := map[string]int{"a": 1, "b": 2, "c": 3}
	for key, value := range mymap {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}

	fmt.Println("For Loop - with a continue statement:")
	// 6. For loop with a continue statement
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			continue // Skip even values
		}
		fmt.Println(i) // Print odd values
	}
}
