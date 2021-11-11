package models

import (
	"blog/vendors/helpers"

	"gorm.io/gorm"

	strip "github.com/grokify/html-strip-tags-go"
)

type Article struct {
	Model
	Title       string `json:"title"`
	Slug        string `json:"slug"`
	Author      string `json:"author"`
	ContentType uint   `json:"content_type"`
	Markdown    string `gorm:"->:false;<-:create" json:"markdown"`
	Html        string `json:"html"`
	Description string `json:"description"`
	Keywords    string `json:"keywords"`
	IsTop       int    `json:"is_top"`
	IsShow      int    `json:"is_show"`
	Views       int    `json:"views"`
	Order       int    `json:"order"`
	CategoryId  int    `json:"category_id"`
	ShotHtml    string `gorm:"-"`

	Category Category `gorm:"foreignKey:CategoryId" json:"category"`
	Tags     []*Tag   `gorm:"many2many:article_tag;" json:"tags"`
}

func (article *Article) AfterFind(tx *gorm.DB) (err error) {
	if article.ShotHtml == "" {
		article.ShotHtml = article.getShortHtml()
	}
	return
}

func (article *Article) getShortHtml() string {
	striped := strip.StripTags(article.Html)
	return helpers.Substr(striped, 0, 200, "")
}
