package basics

import (
	"fmt"
	"os"
)

func ErrorHandling() {
	fmt.Println("Error Handling:")
	file, err := os.Open("nonexistent_file.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	file.Close()
}
