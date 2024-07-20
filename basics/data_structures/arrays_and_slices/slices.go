package arrays_and_slices

import "fmt"

func Slices() {
	slice := []int{1, 2, 3, 4, 5}
	slice = append(slice, 6)
	fmt.Println("Slices:")
	fmt.Println(slice)
}
