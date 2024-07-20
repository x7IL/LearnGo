package additional_concepts

import "fmt"

func VariableScope() {
	globalVar := "I'm global"

	fmt.Println("Variable Scope:")
	fmt.Println(globalVar)

	func() {
		localVar := "I'm local"
		fmt.Println(localVar)
		fmt.Println(globalVar)
	}()
	// fmt.Println(localVar) // This will cause an error
}
