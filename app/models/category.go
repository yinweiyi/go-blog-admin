package models

type Category struct {
	Model
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	Order       int    `json:"order"`
}
