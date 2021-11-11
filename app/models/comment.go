package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	Model

	ParentId        int    `json:"parent_id"`
	Content         string `json:"content"`
	Avatar          string `json:"avatar"`
	Nickname        string `json:"nickname"`
	Email           string `json:"email"`
	CommentableId   int    `json:"commentable_id"`
	CommentableType string `json:"commentable_type"`
	Ip              string `json:"ip"`
	IsAudited       int    `json:"is_audited"`
	IsRead          int    `json:"is_read"`
	IsAdminReply    int    `json:"is_admin_reply"`
	TopId           int    `json:"top_id"`

	Children []Comment `gorm:"-" json:"children"`
}

func (u *Comment) AfterFind(tx *gorm.DB) (err error) {
	//u.Avatar = "/assets/" + u.Avatar
	return
}
