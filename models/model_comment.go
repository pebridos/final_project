package models

type Comment struct {
	GormModel
	UserID  uint
	PhotoID uint
	Message string `gorm:"size:250" json:"message"`
	User    User
	Photo   Photo
}
