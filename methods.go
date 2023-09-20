package main

import (
	"bufio"
	"encoding/json"
	"errors"
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
			m.ParseStrings(lines, "lora")
		}
	}
}

func (m *Models) printEach() {
	var allModels []Printables

	allModels = m.appendEach()

	log.Print(printModels(allModels))
}

func (m *Models) jsonAll() (byteArray []byte, err error) {
	return json.MarshalIndent(m, "", "    ")
}

func (m *Models) jsonAllAndSave() error {
	var dataTypes = map[string]any{
		"loras":       m.Loras,
		"checkpoints": m.Checkpoints,
		"vaes":        m.Vaes,
		"embeddings":  m.Embeddings,
	}

	for dataType, data := range dataTypes {
		switch dataSlice := data.(type) {
		case []*Lora:
			for _, item := range dataSlice {
				bytes, err := item.ToJSON()
				if err != nil {
					return err
				}
				fileName := dataType + ".json"
				SaveJsonToFile(fileName, bytes)
			}
		case []*Checkpoint:
			for _, item := range dataSlice {
				bytes, err := item.ToJSON()
				if err != nil {
					return err
				}
				fileName := dataType + ".json"
				SaveJsonToFile(fileName, bytes)
			}
		case []*Vae:
			for _, item := range dataSlice {
				bytes, err := item.ToJSON()
				if err != nil {
					return err
				}
				fileName := dataType + ".json"
				SaveJsonToFile(fileName, bytes)
			}
		case []*Embedding:
			for _, item := range dataSlice {
				bytes, err := item.ToJSON()
				if err != nil {
					return err
				}
				fileName := dataType + ".json"
				SaveJsonToFile(fileName, bytes)
			}
		default:
			return errors.New("unknown datatype encountered")
		}
	}

	return nil
}

func (m *Models) appendEach() (print []Printables) {
	var allModels []Printables

	for _, lora := range m.Loras {
		m.Loras = append(m.Loras, lora)
		allModels = append(allModels, lora)
	}
	for _, checkpoint := range m.Checkpoints {
		m.Checkpoints = append(m.Checkpoints, checkpoint)
		allModels = append(allModels, checkpoint)
	}
	for _, vae := range m.Vaes {
		m.Vaes = append(m.Vaes, vae)
		allModels = append(allModels, vae)
	}
	for _, embedding := range m.Embeddings {
		m.Embeddings = append(m.Embeddings, embedding)
		allModels = append(allModels, embedding)
	}

	return allModels
}

func printModels[T Printables](models []T) string {
	var toPrint []string
	for _, model := range models {
		toPrint = append(toPrint, model.SPrint())
	}
	return strings.Join(toPrint, "\n")
}

func bytesJsonModels(m *Printables) (bytes []byte, err error) {
	return (*m).ToJSON()
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

		m.ParseStrings(line, currentSection)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func (m *Models) ParseStrings(text string, option string) {
	// example: "artist/artistLoRa.safetensors;"
	// example: "character/HeroCharacter.ckpt;"
	compile, err := regexp.Compile(`(?P<folder>\w+[\\/])?(?P<filename>.*)(?P<extension>\.(?:safetensors|(?:ck)?pt));?`)
	if err != nil {
		return
	}

	match := compile.FindStringSubmatch(text)
	if match == nil {
		return
	}

	result := make(map[string]string)
	for i, name := range compile.SubexpNames() {
		if i != 0 && name != "" && i < len(match) {
			result[name] = match[i]
		}
	}

	switch option {
	case "loras":
		m.Loras = append(m.Loras, &Lora{Folder: result["folder"], Filename: result["filename"], Extension: result["extension"]})
	case "checkpoints":
		m.Checkpoints = append(m.Checkpoints, &Checkpoint{Folder: result["folder"], Filename: result["filename"], Extension: result["extension"]})
	case "vaes":
		m.Vaes = append(m.Vaes, &Vae{Folder: result["folder"], Filename: result["filename"], Extension: result["extension"]})
	case "embeddings":
		m.Embeddings = append(m.Embeddings, &Embedding{Folder: result["folder"], Filename: result["filename"], Extension: result["extension"]})
	}
}

func SaveJsonToFile(fileName string, json []byte) {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		if closeErr := file.Close(); closeErr != nil {
			log.Fatal(closeErr)
		}
	}(file)

	_, err = file.Write(json)
	if err != nil {
		log.Fatal(err)
	}
}
