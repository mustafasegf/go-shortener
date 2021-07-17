package repository

import "gorm.io/gorm"

type Link struct {
	Db *gorm.DB
}

func NewLinkRepo(db *gorm.DB) *Link {
	return &Link{
		Db: db,
	}
}