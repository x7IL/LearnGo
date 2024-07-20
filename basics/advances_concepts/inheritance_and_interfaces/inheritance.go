package inheritance_and_interfaces

import "fmt"

type Animal struct {
	Name string
}

func (a Animal) Speak() string {
	return "Some sound"
}

type Dog struct {
	Animal
	Breed string
}

func Inheritance() {
	dog := Dog{Animal: Animal{Name: "Rex"}, Breed: "Labrador"}
	fmt.Println("Inheritance (Embedded Structs):")
	fmt.Println(dog.Name)
	fmt.Println(dog.Breed)
	fmt.Println(dog.Speak())
}
