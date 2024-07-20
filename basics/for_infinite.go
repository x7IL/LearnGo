package basics

import "fmt"

func ForInfinite() {
	fmt.Println("For Infinite Loop:")
	i := 0
	for {
		fmt.Println(i)
		i++
		if i > 5 {
			break // Exit the infinite loop after printing 0 to 5
		}
	}
}
