package models

type About struct {
	Model

	Title       string `json:"title"`
	ContentType int    `json:"content_type"`
	Markdown    string `json:"markdown"`
	Html        string `json:"html"`
	Order       int    `json:"order"`
	IsEnable    int    `json:"is_enable"`
}
