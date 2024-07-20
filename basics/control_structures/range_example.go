package control_structures

import "fmt"

// Function demonstrating various uses of the range keyword
func RangeExample() {
	// Iterating over a slice
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Range over slice:")
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// Iterating over a map
	fruits := map[string]string{"apple": "red", "banana": "yellow"}
	fmt.Println("Range over map:")
	for key, value := range fruits {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}

	// Iterating over a string
	text := "hello"
	fmt.Println("Range over string:")
	for index, runeValue := range text {
		fmt.Printf("Index: %d, Rune: %c\n", index, runeValue)
	}

	// Using range with only values in a slice
	fmt.Println("Range over slice (only values):")
	for _, value := range numbers {
		fmt.Println("Value:", value)
	}

	// Using range with only keys in a map
	fmt.Println("Range over map (only keys):")
	for key := range fruits {
		fmt.Println("Key:", key)
	}

	// Skipping values in a string
	fmt.Println("Range over string (skipping values):")
	for _, runeValue := range text {
		fmt.Printf("Rune: %c\n", runeValue) // Only printing rune values
	}

	// Using range with multiple variables
	fmt.Println("Range over slice (multiple variables):")
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
}
