package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Please provide file path")
		os.Exit(1)
	}

	filePath := os.Args[1]
	fmt.Println("File path: ", filePath)

	_, err := os.Stat(filePath)
	if err != nil {
		fmt.Println("Please provide correct file path")
		fmt.Println(err)
		os.Exit(1)
	}

	file, fileError := os.ReadFile(filePath)

	if fileError != nil {
		fmt.Println("Read file fileError")
		fmt.Println(fileError)
		os.Exit(1)
	}

	fmt.Println(string(file))
}
