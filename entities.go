package main

type Models struct {
	Loras       []*Lora       `json:"loras,omitempty"`
	Checkpoints []*Checkpoint `json:"checkpoints,omitempty"`
	Vaes        []*Vae        `json:"vaes,omitempty"`
	Embeddings  []*Embedding  `json:"embeddings,omitempty"`
}

type Lora struct {
	Folder    string `json:"folder,omitempty"`
	Filename  string `json:"filename"`
	Extension string `json:"extension"`
}

type Checkpoint struct {
	Folder    string `json:"folder,omitempty"`
	Filename  string `json:"filename"`
	Extension string `json:"extension"`
}

type Vae struct {
	Folder    string `json:"folder,omitempty"`
	Filename  string `json:"filename"`
	Extension string `json:"extension"`
}

type Embedding struct {
	Folder    string `json:"folder,omitempty"`
	Filename  string `json:"filename"`
	Extension string `json:"extension"`
}
