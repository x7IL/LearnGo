package file_operations

import (
	"fmt"
	"os"
)

func FileOperations() {
	// Create a new file
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	file.Close()

	// Rename the file
	err = os.Rename("example.txt", "new_example.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Delete the file
	err = os.Remove("new_example.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("File Operations: Created, renamed, and deleted file.")
}
