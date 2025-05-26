package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Title       string `json:"title,omitempty"`
	IsCompleted bool   `json:"iscompleted"`
}
