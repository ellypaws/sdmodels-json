package main

import (
	"encoding/json"
	"fmt"
)

type Models struct {
	Loras       []*Lora       `json:"loras,omitempty"`
	Checkpoints []*Checkpoint `json:"checkpoints,omitempty"`
	Vaes        []*Vae        `json:"vaes,omitempty"`
	Embeddings  []*Embedding  `json:"embeddings,omitempty"`
}

type Printables interface {
	SPrint() string
	ToJSON() ([]byte, error)
}

type Lora struct {
	Folder    string `json:"folder,omitempty"`
	Filename  string `json:"filename"`
	Extension string `json:"extension"`
}

func (l *Lora) ToJSON() ([]byte, error) {
	return json.MarshalIndent(l, "", "    ")
}

func (l *Lora) SPrint() string {
	return fmt.Sprintf("\n[Lora]\nFolder: %s\nFilename: %s\nExtension: %s\n", l.Folder, l.Filename, l.Extension)
}

type Checkpoint struct {
	Folder    string `json:"folder,omitempty"`
	Filename  string `json:"filename"`
	Extension string `json:"extension"`
}

func (c Checkpoint) ToJSON() ([]byte, error) {
	return json.MarshalIndent(c, "", "    ")
}

func (c Checkpoint) SPrint() string {
	return fmt.Sprintf("\n[Checkpoint]\nFolder: %s\nFilename: %s\nExtension: %s\n", c.Folder, c.Filename, c.Extension)
}

type Vae struct {
	Folder    string `json:"folder,omitempty"`
	Filename  string `json:"filename"`
	Extension string `json:"extension"`
}

func (v Vae) ToJSON() ([]byte, error) {
	return json.MarshalIndent(v, "", "    ")
}

func (v Vae) SPrint() string {
	return fmt.Sprintf("\n[Vae]\nFolder: %s\nFilename: %s\nExtension: %s\n", v.Folder, v.Filename, v.Extension)
}

type Embedding struct {
	Folder    string `json:"folder,omitempty"`
	Filename  string `json:"filename"`
	Extension string `json:"extension"`
}

func (e Embedding) ToJSON() ([]byte, error) {
	return json.MarshalIndent(e, "", "    ")
}

func (e Embedding) SPrint() string {
	return fmt.Sprintf("\n[Embedding]\nFolder: %s\nFilename: %s\nExtension: %s\n", e.Folder, e.Filename, e.Extension)
}
