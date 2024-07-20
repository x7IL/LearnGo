package control_structures

import "fmt"

func WhileLoop() {
	count := 0
	fmt.Println("While Loop:")
	for count < 5 {
		fmt.Println(count)
		count++
	}
}
