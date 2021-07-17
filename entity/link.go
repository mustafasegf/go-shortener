package entity

import "gorm.io/gorm"

type CreateLinkRequest struct {
	LongUrl  string `json:"long_url" binding:"required,url"`
	ShortUrl string `json:"short_url" binding:"required"`
}

type LinkModel struct {
	gorm.Model
	LongUrl  string `gorm:"column:long_url;type:varchar(2048)"`
	ShortUrl string `gorm:"column:short_url;type:varchar(255)"`
}
