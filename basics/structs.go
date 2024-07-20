package basics

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
