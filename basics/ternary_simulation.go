package basics

import "fmt"

func ternary(condition bool, trueValue, falseValue string) string {
	if condition {
		return trueValue
	}
	return falseValue
}

func TernarySimulation() {
	age := 18
	result := ternary(age >= 18, "Adult", "Minor")
	fmt.Println("Ternary Simulation:")
	fmt.Println("You are:", result)
}
