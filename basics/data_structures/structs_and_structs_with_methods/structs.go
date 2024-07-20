package structs_and_structs_with_methods

import "fmt"

type Person struct {
	Name string
	Age  int
}

func Structs() {
	p := Person{Name: "Alice", Age: 30}
	fmt.Println("Structs:")
	fmt.Println(p.Name, p.Age)
}
