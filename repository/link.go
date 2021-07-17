package repository

import (
	"github.com/mustafasegf/go-shortener/entity"
	"gorm.io/gorm"
)


type Link struct {
	Db *gorm.DB
}

func NewLinkRepo(db *gorm.DB) *Link {
	return &Link{
		Db: db,
	}
}

func (r *Link) GetLinkByURL(ShortUrl string) (entity entity.CreateLinkRequest, err error) {
	query := r.Db.Table("link").
		Where("short_url = ?", ShortUrl).
		First(&entity)
	
	err = query.Error
	return
}

func (r *Link) InsertURL(shortUrl, longUrl string) (err error) {
	model := entity.LinkModel{
		LongUrl: longUrl,
		ShortUrl: shortUrl,
	}

	query := r.Db.Table("link").
		Create(&model)
	
	err = query.Error
	return
}
