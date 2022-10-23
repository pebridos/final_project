package models

type SocialMedia struct {
	GormModel
	UserID         uint
	Name           string `gorm:"size:100" json:"socialmedia_name"`
	Caption        string `gorm:"size:250" json:"socialmedia_caption"`
	SocialMediaURL string `gorm:"size:250" json:"socialmedia_url"`
	User           User
}
