package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"
)

func Create() *Models {
	return &Models{}
}

func (m *Models) CreateLora() {
	m.Loras = []*Lora{}
}

func (m *Models) CreateCheckpoint() {
	m.Checkpoints = []*Checkpoint{}
}

func (m *Models) CreateVae() {
	m.Vaes = []*Vae{}
}

func (m *Models) CreateEmbedding() {
	m.Embeddings = []*Embedding{}
}

func (m *Models) CreateAll() {
	m.CreateLora()
	m.CreateCheckpoint()
	m.CreateVae()
	m.CreateEmbedding()
}

func (m *Models) ReadLoraFromFile(fileName string) {
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0644)
	if err != nil {
		return
	}
	defer func(file *os.File) { _ = file.Close() }(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//log.Println(scanner.Text())
		lines := scanner.Text()
		if lines != "" {
			m.StringToLora(lines)
		}
	}
}

func (m *Models) printEach() {
	var allModels []Printables

	for _, v := range m.Loras {
		allModels = append(allModels, v)
	}
	for _, v := range m.Checkpoints {
		allModels = append(allModels, v)
	}
	for _, v := range m.Vaes {
		allModels = append(allModels, v)
	}
	for _, v := range m.Embeddings {
		allModels = append(allModels, v)
	}

	log.Print(printModels(allModels))
}

func printModels[T Printables](models []T) string {
	var toPrint []string
	for _, model := range models {
		toPrint = append(toPrint, model.SPrint())
	}
	return strings.Join(toPrint, "\n")
}

func (m *Models) ReadFromFileAndSort(fileName string) {
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	scanner := bufio.NewScanner(file)
	var currentSection string
	headerRegex := regexp.MustCompile(`_{2,}(?P<header>[a-zA-Z\s]+?)_{2,}`)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		match := headerRegex.FindStringSubmatch(line)
		if len(match) > 1 {
			currentSection = strings.ToLower(strings.TrimSpace(match[1]))
			continue
		}

		switch currentSection {
		case "loras":
			m.StringToLora(line)
		case "checkpoints":
			//m.StringToCheckpoint(line)
		case "vaes":
			m.StringToVae(line)
		case "schedulers", "samplers":
			// Add cases based on suitability
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func (m *Models) StringToLora(input string) {
	// example: "artist/artistLoRa.safetensors;"
	// example: "character/HeroCharacter.ckpt;"
	compile, err := regexp.Compile(`(?P<folder>\w+[\\/])?(?P<filename>.*)(?P<extension>\.(?:safetensors|(?:ck)?pt));?`)
	if err != nil {
		return
	}

	match := compile.FindStringSubmatch(input)
	if match == nil {
		return
	}

	result := make(map[string]string)
	for i, name := range compile.SubexpNames() {
		if i != 0 && name != "" && i < len(match) {
			result[name] = match[i]
		}
	}

	lora := &Lora{
		Folder:    result["folder"],
		Filename:  result["filename"],
		Extension: result["extension"],
	}

	// check if Lora is nil, use CreateLora() if it is
	if m.Loras == nil {
		m.CreateLora()
	}
	m.Loras = append(m.Loras, lora)

	if lora.Filename == "" {
		return
	}

	return
}

func (m *Models) StringToVae(input string) {
	compile, err := regexp.Compile(`(?P<folder>\w+[\\/])?(?P<filename>.*?)(?P<extension>\.(?:ck)?pt);?`)
	if err != nil {
		return
	}

	match := compile.FindStringSubmatch(input)
	if match == nil {
		return
	}

	result := make(map[string]string)
	for i, name := range compile.SubexpNames() {
		if i != 0 && name != "" && i < len(match) {
			result[name] = match[i]
		}
	}

	vae := &Vae{
		Folder:    result["folder"],
		Filename:  result["filename"],
		Extension: result["extension"],
	}

	// check if Vae is nil, use CreateVae() if it is
	if m.Vaes == nil {
		m.CreateVae()
	}
	m.Vaes = append(m.Vaes, vae)

	return
}
