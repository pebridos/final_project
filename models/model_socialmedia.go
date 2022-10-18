package models

import "time"

type SocialMedia struct {
	ID             uint      `gorm:"primayKey" json:"socialmedia_id"`
	UserID         uint      `json:"user_id"`
	Name           string    `gorm:"size:100" json:"socialmedia_name"`
	Caption        string    `gorm:"size:250" json:"socialmedia_caption"`
	SocialMediaURL string    `gorm:"size:250" json:"socialmedia_url"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
