package additional_concepts

import "fmt"

// ContinueBreak demonstrates the use of 'continue' and 'break' statements in loops.
func ContinueBreak() {
	fmt.Println("Continue and Break Example:")

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue // Skip even numbers
		}
		if i > 7 {
			break // Exit the loop if the number is greater than 7
		}
		fmt.Println(i)
	}
}
