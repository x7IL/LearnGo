package main

import (
	"LearnGo/basics"
	"LearnGo/basics/additional_concepts"
	"LearnGo/basics/advances_concepts/error_handling"
	"LearnGo/basics/advances_concepts/file_operations"
	"LearnGo/basics/advances_concepts/inheritance_and_interfaces"
	"LearnGo/basics/control_structures"
	"LearnGo/basics/data_structures/arrays_and_slices"
	"LearnGo/basics/data_structures/maps"
	"LearnGo/basics/data_structures/structs_and_structs_with_methods"
	"LearnGo/basics/expressions_and_operators"
	"LearnGo/basics/functions"
	"LearnGo/basics/variable_data_and_types"
)

func main() {
	// Basic Concepts

	// 1. Hello World
	basics.HelloWorld()

	// 2. Variables and Data Types
	variable_data_and_types.Variables()
	variable_data_and_types.DataTypes()

	// 3. Expressions and Operators
	expressions_and_operators.Expressions()
	expressions_and_operators.Operators()        // Includes Arithmetic Operators, etc.
	expressions_and_operators.OperatorsLogical() // Logical Operators

	// 4. Control Structures
	control_structures.IfElse()
	control_structures.SwitchExample()
	control_structures.ForLoop()
	control_structures.ForInfinite()
	control_structures.WhileLoop() // Ensure that you have this function if it’s used.
	control_structures.RangeExample()

	// 5. Functions
	functions.Functions()
	functions.TernarySimulation() // Simulating ternary operator

	// Data Structures

	// 1. Arrays and Slices
	arrays_and_slices.Arrays()
	arrays_and_slices.Slices()

	// 2. Maps
	maps.Maps()

	// 3. Structs and Structs with Methods
	structs_and_structs_with_methods.Structs()
	structs_and_structs_with_methods.StructsWithMethods()

	// Advanced Concepts

	// 1. Error Handling
	error_handling.ErrorHandling()

	// 2. File Operations
	file_operations.ReadingFiles()
	file_operations.WritingFiles()
	file_operations.FileOperations()

	// 3. Inheritance and Interfaces
	inheritance_and_interfaces.Inheritance()
	inheritance_and_interfaces.Interfaces()

	// Additional Concepts
	additional_concepts.VariableScope()
	additional_concepts.Closures()
	additional_concepts.DeferExample()
	additional_concepts.ContinueBreak()
}
