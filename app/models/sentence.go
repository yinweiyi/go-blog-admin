package models

type Sentence struct {
	Model
	Author      string `json:"author"`
	Content     string `json:"content"`
	Translation string `json:"translation"`
}
