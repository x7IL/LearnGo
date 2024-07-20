package main

import "LearnGo/basics"

func main() {
	// Basic Concepts

	// 1. Hello World
	basics.HelloWorld()

	// 2. Variables and Data Types
	basics.Variables()
	basics.DataTypes()

	// 3. Expressions and Operators
	basics.Expressions()
	basics.Operators()        // Includes Arithmetic Operators, etc.
	basics.OperatorsLogical() // Logical Operators

	// 4. Control Structures
	basics.IfElse()
	basics.SwitchExample()
	basics.ForLoop()
	basics.ForInfinite()
	basics.WhileLoop() // Ensure that you have this function if it’s used.
	basics.RangeExample()

	// 5. Functions
	basics.Functions()
	basics.TernarySimulation() // Simulating ternary operator

	// Data Structures

	// 1. Arrays and Slices
	basics.Arrays()
	basics.Slices()

	// 2. Maps
	basics.Maps()

	// 3. Structs and Structs with Methods
	basics.Structs()
	basics.StructsWithMethods()

	// Advanced Concepts

	// 1. Error Handling
	basics.ErrorHandling()

	// 2. File Operations
	basics.ReadingFiles()
	basics.WritingFiles()
	basics.FileOperations()

	// 3. Inheritance and Interfaces
	basics.Inheritance()
	basics.Interfaces()

	// Additional Concepts
	basics.VariableScope()
	basics.Closures()
	basics.DeferExample()
	basics.ContinueBreak()
}
