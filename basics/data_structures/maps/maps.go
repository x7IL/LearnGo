package maps

import "fmt"

func Maps() {
	myMap := make(map[string]int)
	myMap["age"] = 30
	fmt.Println("Maps:")
	fmt.Println(myMap["age"])
}
