package models

type Todo struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Title       string `json:"title,omitempty"`
	IsCompleted bool   `json:"iscompleted"`
}
