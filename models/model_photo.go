package models

import "time"

type Photo struct {
	ID        uint      `gorm:"primayKey" json:"photo_id"`
	UserID    uint      `json:"user_id"`
	Title     string    `gorm:"size:100" json:"photo_title"`
	Caption   string    `gorm:"size:250" json:"photo_caption"`
	PhotoURL  string    `gorm:"size:250" json:"photo_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
