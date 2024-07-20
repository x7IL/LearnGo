package additional_concepts

import "fmt"

func DeferExample() {
	fmt.Println("Defer Example:")

	// Deferred function will run after the surrounding function finishes
	defer fmt.Println("Deferred: This will run after the function completes")
	fmt.Println("This will run first")
}
