package models

type Config struct {
	Model

	Title       string `json:"title"`
	SubTitle    string `json:"sub_title"`
	Keywords    string `json:"keywords"`
	Icp         string `json:"icp"`
	Author      string `json:"author"`
	Description string `json:"description"`
}
