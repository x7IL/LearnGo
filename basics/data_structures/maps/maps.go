package maps

import "fmt"

// Function to create and initialize a map
func createMap() map[string]int {
	return map[string]int{
		"Alice": 25,
		"Bob":   30,
		"Carol": 35,
	}
}

// Function to print a map
func printMap(m map[string]int) {
	fmt.Println("Map contents:")
	for key, value := range m {
		fmt.Printf("%s: %d\n", key, value)
	}
}

// Function to add or update a key-value pair in a map
func addOrUpdateMap(m map[string]int, key string, value int) {
	m[key] = value
}

// Function to delete a key-value pair from a map
func deleteFromMap(m map[string]int, key string) {
	delete(m, key)
}

// Function to check if a key exists in a map
func keyExists(m map[string]int, key string) bool {
	_, exists := m[key]
	return exists
}

// Function to demonstrate various map operations
func Maps() {
	// Create and print a map
	myMap := createMap()
	fmt.Println("Original Map:")
	printMap(myMap)

	// Add or update key-value pairs
	addOrUpdateMap(myMap, "Dave", 40)
	fmt.Println("Map after adding Dave:")
	printMap(myMap)

	// Delete a key-value pair
	deleteFromMap(myMap, "Bob")
	fmt.Println("Map after deleting Bob:")
	printMap(myMap)

	// Check if a key exists
	if keyExists(myMap, "Alice") {
		fmt.Println("Alice exists in the map.")
	} else {
		fmt.Println("Alice does not exist in the map.")
	}

	// Access a specific value
	fmt.Println("Value associated with key 'Alice':", myMap["Alice"])
}
