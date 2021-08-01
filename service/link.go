package service

import (
	"log"
	"net/url"

	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/copier"
	"github.com/mustafasegf/go-shortener/entity"
	"github.com/mustafasegf/go-shortener/repository"
)

type Link struct {
	repo *repository.Link
}

func NewLinkService(repo *repository.Link) *Link {
	return &Link{
		repo: repo,
	}
}

func (s *Link) GetLinkByURL(shortUrl string) (result *entity.CreateLinkRequest, err error) {
	result = &entity.CreateLinkRequest{}
	longURL, err := s.repo.RedisGetLinkByURL(shortUrl)
	if err != redis.Nil {
		result.LongUrl = longURL
		result.ShortUrl = shortUrl
		err = nil
		return
	}
	data, err := s.repo.GetLinkByURL(shortUrl)
	if err != nil {
		return
	}
	err = s.repo.RedisSetURL(shortUrl, data.LongUrl)
	if err != redis.Nil {
		log.Print(err)
		err = nil
	}
	err = copier.Copy(&result, data)

	return
}

func (s *Link) InsertURL(req entity.CreateLinkRequest) (err error) {
	err = s.repo.InsertURL(req.ShortUrl, req.LongUrl)
	s.repo.RedisSetURL(req.ShortUrl, req.LongUrl)
	return
}

func (s *Link) CheckURL(longUrl string) bool {
	if longUrl == "http://" {
		return false
	}
	_, err := url.ParseRequestURI(longUrl)
	if err != nil {
		return false
	}
	return true
}
