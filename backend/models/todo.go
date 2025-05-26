package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Title       string `json:"title"`
	IsCompleted bool   `json:"iscompleted"`
}
type Input struct {
	Title       string `json:"title" binding:"required"`
	IsCompleted bool   `json:"iscompleted"`
}
