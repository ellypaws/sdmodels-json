package main

import (
	"fmt"
	"os"
)

func main() {
	var filename string
	// if loras.txt exists in current directory, use that
	if _, err := os.Stat("loras.txt"); err == nil {
		filename = "loras.txt"
	} else {
		fmt.Print("Enter the filename: ")
		_, _ = fmt.Scan(&filename)
	}
	models := Create()
	//models.CreateLora()
	//models.ReadLoraFromFile(filename)
	//models.printEach()
	models.ReadFromFileAndSort(filename)
	models.printEach()
	bytes, err := models.jsonEach()
	if err != nil {
		return
	}
	fmt.Println(string(bytes))
	SaveJsonToFile("models.json", bytes)
}
