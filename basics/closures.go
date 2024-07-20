package basics

import "fmt"

func Closures() {
	fmt.Println("Closures Example:")

	// Function that returns another function (closure)
	counter := func() func() int {
		i := 0
		return func() int {
			i++
			return i
		}
	}()

	// Demonstrating the closure
	fmt.Println(counter()) // Output: 1
	fmt.Println(counter()) // Output: 2
	fmt.Println(counter()) // Output: 3
}
