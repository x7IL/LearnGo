package basics

import "fmt"

func Operators() {
	a := 5
	b := 2
	sum := a + b
	difference := a - b
	product := a * b
	quotient := a / b

	fmt.Println("Operators:")
	fmt.Println("sum:", sum)
	fmt.Println("difference:", difference)
	fmt.Println("product:", product)
	fmt.Println("quotient:", quotient)

	fmt.Println("Comparison Operators:")
	fmt.Println("a == b:", a == b)
	fmt.Println("a > b:", a > b)
}
