package expressions_and_operators

import "fmt"

func Expressions() {
	x := 10
	y := 5

	// Arithmetic Operators
	add := x + y // Addition
	sub := x - y // Subtraction
	mul := x * y // Multiplication
	div := x / y // Division
	mod := x % y // Modulus
	inc := x + 1 // Increment (using addition)
	dec := y - 1 // Decrement (using subtraction)

	// Comparison Operators
	eq := x == y // Equal to
	ne := x != y // Not equal to
	lr := x < y  // Less than
	gr := x > y  // Greater than
	le := x <= y // Less than or equal to
	ge := x >= y // Greater than or equal to

	// Logical Operators
	and := (x > 5 && y < 10) // Logical AND
	or := (x > 5 || y < 10)  // Logical OR
	not := !(x > 5)          // Logical NOT

	// Bitwise Operators
	bitAnd := x & y         // Bitwise AND
	bitOr := x | y          // Bitwise OR
	bitXor := x ^ y         // Bitwise XOR
	bitNot := ^x            // Bitwise NOT
	bitShiftLeft := x << 2  // Bitwise left shift
	bitShiftRight := x >> 1 // Bitwise right shift

	fmt.Println("Expressions:")
	fmt.Println("Arithmetic Operators:")
	fmt.Println("Addition:", add)
	fmt.Println("Subtraction:", sub)
	fmt.Println("Multiplication:", mul)
	fmt.Println("Division:", div)
	fmt.Println("Modulus:", mod)
	fmt.Println("Increment (x + 1):", inc)
	fmt.Println("Decrement (y - 1):", dec)

	fmt.Println("\nComparison Operators:")
	fmt.Println("Equal to:", eq)
	fmt.Println("Not equal to:", ne)
	fmt.Println("Less than:", lr)
	fmt.Println("Greater than:", gr)
	fmt.Println("Less than or equal to:", le)
	fmt.Println("Greater than or equal to:", ge)

	fmt.Println("\nLogical Operators:")
	fmt.Println("Logical AND (x > 5 && y < 10):", and)
	fmt.Println("Logical OR (x > 5 || y < 10):", or)
	fmt.Println("Logical NOT (!(x > 5)):", not)

	fmt.Println("\nBitwise Operators:")
	fmt.Println("Bitwise AND (x & y):", bitAnd)
	fmt.Println("Bitwise OR (x | y):", bitOr)
	fmt.Println("Bitwise XOR (x ^ y):", bitXor)
	fmt.Println("Bitwise NOT (^x):", bitNot)
	fmt.Println("Bitwise left shift (x << 2):", bitShiftLeft)
	fmt.Println("Bitwise right shift (x >> 1):", bitShiftRight)
}
