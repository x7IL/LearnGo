package expressions_and_operators

import "fmt"

func OperatorsLogical() {
	a := true
	b := false

	fmt.Println("Logical Operators:")
	fmt.Println("a && b:", a && b) // AND
	fmt.Println("a || b:", a || b) // OR
	fmt.Println("!a:", !a)         // NOT
}
