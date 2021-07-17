package service

import "github.com/mustafasegf/go-shortener/repository"

type Link struct {
	repo *repository.Link
}

func NewLinkService(repo *repository.Link) *Link {
	return &Link{
		repo: repo,
	}
}