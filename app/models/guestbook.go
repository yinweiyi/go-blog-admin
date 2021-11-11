package models

type Guestbook struct {
	Model

	ContentType int    `json:"content_type"`
	Markdown    string `json:"markdown"`
	Html        string `json:"html"`
	CanComment  int    `json:"can_comment"`
}

func (g Guestbook) TableName() string {
	return "guestbook"
}
