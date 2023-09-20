package main

import (
	"fmt"
	"os"
)

func main() {
	var filename string
	// if models.txt exists in current directory, use that
	if _, err := os.Stat("models.txt"); err == nil {
		filename = "models.txt"
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
	bytes, err := models.jsonAll()
	if err != nil {
		return
	}
	fmt.Println(string(bytes))
	SaveJsonToFile("models.json", bytes)

	err = models.jsonAllAndSave()
	if err != nil {
		return
	}
}
