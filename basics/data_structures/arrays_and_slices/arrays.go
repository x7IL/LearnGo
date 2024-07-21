package arrays_and_slices

import "fmt"

// Function to define and initialize a basic array of integers
func defineIntArray() [5]int {
	arr := [5]int{1, 2, 3, 4, 5}
	return arr
}

// Function to print a basic array of integers
func printIntArray(arr [5]int) {
	fmt.Println("Integer Array:")
	for _, value := range arr {
		fmt.Println(value)
	}
}

// Function to define and initialize an array of strings
func defineStringArray() [3]string {
	arr := [3]string{"Go", "is", "fun"}
	return arr
}

// Function to print an array of strings
func printStringArray(arr [3]string) {
	fmt.Println("String Array:")
	for _, value := range arr {
		fmt.Println(value)
	}
}

// Function to define and initialize a two-dimensional array
func define2DArray() [2][3]int {
	arr := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	return arr
}

// Function to sum the elements of a two-dimensional array
func sum2DArray(arr [2][3]int) int {
	total := 0
	for _, row := range arr {
		for _, value := range row {
			total += value
		}
	}
	return total
}

// Function to print a two-dimensional array
func print2DArray(arr [2][3]int) {
	fmt.Println("2D Array:")
	for _, row := range arr {
		for _, value := range row {
			fmt.Print(value, " ")
		}
		fmt.Println()
	}
}

// Function to demonstrate the usage of various array operations
func ArrayExamples() {
	// Define and print a basic integer array
	intArr := defineIntArray()
	printIntArray(intArr)

	// Define and print an array of strings
	strArr := defineStringArray()
	printStringArray(strArr)

	// Define a two-dimensional array and print it
	matrix := define2DArray()
	print2DArray(matrix)

	// Calculate and print the sum of elements in the two-dimensional array
	sum := sum2DArray(matrix)
	fmt.Printf("Sum of 2D Array elements: %d\n", sum)
}
