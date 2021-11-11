package models

type FriendshipLink struct {
	Model
	Title       string `json:"title"`
	Link        string `json:"link"`
	Description string `json:"description"`
	IsEnable    int    `json:"is_enable"`
}
