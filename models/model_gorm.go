package models

import "time"

type GormModel struct {
	ID        uint       `gorm:"primayKey" json:"id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
