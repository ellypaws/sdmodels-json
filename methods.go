package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

	return
}

func (m *Models) printEach() {
	for _, lora := range m.Loras {
		fmt.Println(lora)
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

	m.Loras = append(m.Loras, lora)

	if lora.Filename == "" {
		return
	}

	return
}
