package control_structures

import "fmt"

func SwitchExample() {
	day := 2
	fmt.Println("Switch:")
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	default:

		fmt.Println("Invalid day")
	}
}
