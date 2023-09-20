package main

import (
	"fmt"
	"os"
)

func main() {
	var filename string
	currentDirectory := os.Getenv("PWD")
	fmt.Printf("Current directory: %s\n", currentDirectory)
	fmt.Print("Enter the filename: ")
	// if loras.txt exists in current directory, use that
	if _, err := os.Stat("loras.txt"); err == nil {
		filename = "loras.txt"
	} else {
		_, _ = fmt.Scan(&filename)
	}
	models := Create()
	models.CreateLora()
	models.ReadLoraFromFile(filename)
	models.printEach()
	//StringToLora(loraSlice)
}
