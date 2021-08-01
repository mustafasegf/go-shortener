package service

import (
	"net/url"

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
	data, err := s.repo.GetLinkByURL(shortUrl)
	if err != nil {
		return
	}
	result = &entity.CreateLinkRequest{}
	err = copier.Copy(&result, data)

	return
}

func (s *Link) InsertURL(req entity.CreateLinkRequest) (err error) {
	err = s.repo.InsertURL(req.ShortUrl, req.LongUrl)
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
