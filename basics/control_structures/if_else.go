package control_structures

import "fmt"

func IfElse() {
	age := 18
	hasPermission := true

	fmt.Println("If-Else:")
	if age >= 18 && hasPermission {
		fmt.Println("Adult with permission")
	} else if age >= 18 {
		fmt.Println("Adult without permission")
	} else {
		fmt.Println("Minor")
	}
}
