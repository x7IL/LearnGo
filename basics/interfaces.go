package basics

import "fmt"

type Speaker interface {
	Speak() string
}

type Person2 struct {
	Name string
}

func (p Person2) Speak() string {
	return "Hello, my name is " + p.Name
}

func greet2(s Speaker) {
	fmt.Println(s.Speak())
}

func Interfaces() {
	p := Person2{Name: "Alice"}
	fmt.Println("Interfaces:")
	greet2(p)
}
