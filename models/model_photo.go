package models

type Photo struct {
	GormModel
	UserID   uint
	Title    string `gorm:"size:100" json:"photo_title"`
	Caption  string `gorm:"size:250" json:"photo_caption"`
	PhotoURL string `gorm:"size:250" json:"photo_url"`
	User     User
}
