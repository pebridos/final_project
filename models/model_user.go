package models

import "time"

type User struct {
	ID        uint      `gorm:"primayKey" json:"socialmedia_id"`
	Username  string    `gorm:"size:50" json:"username"`
	Email     string    `gorm:"size:40" json:"email"`
	Password  string    `gorm:"size:20" json:"password"`
	Age       int       `json:"age"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
