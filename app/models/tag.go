package models

import (
	"math/rand"
)

type Tag struct {
	Model
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	Order       int    `json:"order"`

	Articles []*Article `gorm:"many2many:article_tag;" json:"articles"`
}

type MinTag struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

//Shuffle 随机
func Shuffle(mins []MinTag) []MinTag {
	values := make([]MinTag, len(mins))
	buf := make([]MinTag, len(mins))
	copy(buf, mins)
	for i := range values {
		j := rand.Intn(len(buf))
		values[i] = buf[j]
		buf = append(buf[0:j], buf[j+1:]...)
	}
	return values
}
