package basics

import "fmt"

func greet(name string) string {
	return "Hello, " + name
}

func add(a int, b int) int {
	return a + b
}

func square(x int) int {
	return x * x
}

func Functions() {
	fmt.Println("Defining Functions:")
	fmt.Println(greet("Alice"))
	fmt.Println("Function Parameters:")
	fmt.Println(add(5, 3))
	fmt.Println("Return Values:")
	fmt.Println(square(4))
}
