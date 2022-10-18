package models

import "time"

type Comment struct {
	ID        uint      `gorm:"primayKey" json:"comment_id"`
	UserID    uint      `json:"user_id"`
	PhotoID   uint      `json:"id_photo"`
	Message   string    `gorm:"size:250" json:"message"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
