package file_operations

import (
	"fmt"
	"os"
)

func WritingFiles() {
	err := os.WriteFile("example.txt", []byte("Hello, Go!"), 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Writing Files: Written to example.txt")
}
