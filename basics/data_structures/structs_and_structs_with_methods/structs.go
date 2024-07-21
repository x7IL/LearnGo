package structs_and_structs_with_methods

import "fmt"

// Definition of a basic structure
type Person struct {
	Name string
	Age  int
}

// Definition of a structure with JSON tags
type PersonWithTags struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Definition of a nested structure
type Address struct {
	Street string
	City   string
}

type PersonWithAddress struct {
	Name    string
	Age     int
	Address Address
}

// Definition of a structure with embedded fields
type PersonWithEmbeddedAddress struct {
	Name    string
	Age     int
	Address // Directly embedded struct
}

// Method associated with Person
func (p Person) Greet() string {
	return "Hello, my name is " + p.Name
}

// Method associated with PersonWithAddress
func (p PersonWithAddress) FullAddress() string {
	return p.Address.Street + ", " + p.Address.City
}

// Method associated with PersonWithEmbeddedAddress
func (p PersonWithEmbeddedAddress) ShortAddress() string {
	return p.Street // Direct access due to embedding
}

func definePerson() Person {
	return Person{Name: "Alice", Age: 30}
}

func definePersonWithTags() PersonWithTags {
	return PersonWithTags{Name: "Bob", Age: 40}
}

func definePersonWithAddress() PersonWithAddress {
	return PersonWithAddress{
		Name: "Charlie",
		Age:  25,
		Address: Address{
			Street: "789 Pine St",
			City:   "Star City",
		},
	}
}

func definePersonWithEmbeddedAddress() PersonWithEmbeddedAddress {
	return PersonWithEmbeddedAddress{
		Name: "Diana",
		Age:  28,
		Address: Address{
			Street: "321 Maple Ave",
			City:   "Central City",
		},
	}
}

func Structures() {
	// Creating and displaying an instance of Person
	p1 := definePerson()
	fmt.Println("Person:")
	fmt.Println("Name:", p1.Name)
	fmt.Println("Age:", p1.Age)
	fmt.Println(p1.Greet())

	// Creating and displaying an instance of PersonWithTags
	p2 := definePersonWithTags()
	fmt.Println("\nPerson with Tags:")
	fmt.Println("Name:", p2.Name)
	fmt.Println("Age:", p2.Age)

	// Creating and displaying an instance of PersonWithAddress
	p3 := definePersonWithAddress()
	fmt.Println("\nPerson with Address:")
	fmt.Println("Name:", p3.Name)
	fmt.Println("Full Address:", p3.FullAddress())

	// Creating and displaying an instance of PersonWithEmbeddedAddress
	p4 := definePersonWithEmbeddedAddress()
	fmt.Println("\nPerson with Embedded Address:")
	fmt.Println("Name:", p4.Name)
	fmt.Println("Short Address:", p4.ShortAddress())
}
