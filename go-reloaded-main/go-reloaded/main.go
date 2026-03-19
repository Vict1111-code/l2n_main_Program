package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run .<inputFile> <outputFile>")
		return
	}

	if os.Args[1] == os.Args[2] {
		fmt.Println("Usage: go run .<inputFile> <outputFile>")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	content, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading file content:", err)
		return
	}

	text := string(content)
	text = ProcessText(text)

	err = os.WriteFile(outputFile, []byte(text), 0644)
	if err != nil {
		fmt.Println("Error writing content to outputFile:", err)
		return
	}
}