package control_structures

import "fmt"

func RangeExample() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Range:")
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
}
