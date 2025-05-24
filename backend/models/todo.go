package models

type Todo struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Title       string `json:"title"`
	IsCompleted bool   `json:"iscompleted"`
}
