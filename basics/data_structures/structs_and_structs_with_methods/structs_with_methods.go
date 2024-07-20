package structs_and_structs_with_methods

import "fmt"

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func StructsWithMethods() {
	rect := Rectangle{Width: 5, Height: 3}
	fmt.Println("Structs with Methods:")
	fmt.Println(rect.Area())
}
