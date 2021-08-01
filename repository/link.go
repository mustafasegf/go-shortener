package repository

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/mustafasegf/go-shortener/entity"
	"gorm.io/gorm"
)

type Link struct {
	Db  *gorm.DB
	Rdb *redis.Client
}

func NewLinkRepo(db *gorm.DB, rdb *redis.Client) *Link {
	return &Link{
		Db:  db,
		Rdb: rdb,
	}
}

func (r *Link) RedisGetLinkByURL(shortUrl string) (longURL string, err error) {
	ctx := context.Background()
	longURL, err = r.Rdb.Get(ctx, shortUrl).Result()
	r.Rdb.Expire(ctx, shortUrl, time.Hour)
	return
}

func (r *Link) RedisSetURL(shortUrl string, longURL string) (err error) {
	ctx := context.Background()
	err = r.Rdb.Set(ctx, shortUrl, longURL, time.Hour).Err()
	return
}

func (r *Link) GetLinkByURL(shortUrl string) (entity entity.LinkModel, err error) {
	query := r.Db.Table("link").
		Where("short_url = ?", shortUrl).
		First(&entity)

	err = query.Error
	return
}

func (r *Link) InsertURL(shortUrl, longUrl string) (err error) {
	model := entity.LinkModel{
		LongUrl:  longUrl,
		ShortUrl: shortUrl,
	}

	query := r.Db.Table("link").
		Create(&model)

	err = query.Error
	return
}
