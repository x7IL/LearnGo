package arrays_and_slices

import "fmt"

// Function to create and initialize a slice
func createSlice() []int {
	return []int{1, 2, 3, 4, 5}
}

// Function to print a slice
func printSlice(slice []int) {
	fmt.Println("Slice elements:")
	for _, value := range slice {
		fmt.Println(value)
	}
}

// Function to append elements to a slice
func appendToSlice(slice []int, elements ...int) []int {
	return append(slice, elements...)
}

// Function to create a sub-slice
func subSlice(slice []int, start, end int) []int {
	return slice[start:end]
}

// Function to demonstrate various slice operations
func Slices() {
	// Create and print a slice
	mySlice := createSlice()
	fmt.Println("Original Slice:")
	printSlice(mySlice)

	// Append elements to the slice
	mySlice = appendToSlice(mySlice, 6, 7, 8)
	fmt.Println("Slice after appending elements:")
	printSlice(mySlice)

	// Create and print a sub-slice
	sub := subSlice(mySlice, 2, 5)
	fmt.Println("Sub-slice (from index 2 to 5):")
	printSlice(sub)

	// Show length and capacity of the slice
	fmt.Printf("Length of slice: %d\n", len(mySlice))
	fmt.Printf("Capacity of slice: %d\n", cap(mySlice))
}
