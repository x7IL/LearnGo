package file_operations

import (
	"fmt"
	"os"
)

func ReadingFiles() {
	data, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Reading Files:")
	fmt.Println(string(data))
}
