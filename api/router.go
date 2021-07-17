package api

import (
	"github.com/mustafasegf/go-shortener/controller"
	"github.com/mustafasegf/go-shortener/service"
	"github.com/mustafasegf/go-shortener/repository"
	"github.com/gin-gonic/gin"
)
type Route struct {
	router *gin.Engine
}


func (s *Server) setupRouter() {
	
	linkRepo := repository.NewLinkRepo(s.db)
	linkSvc := service.NewLinkService(linkRepo)
	linkCtlr := controller.NewLinkController(linkSvc)
	
	api := s.router.Group("/api")
	link := api.Group("/link")

	link.POST("create", linkCtlr.CreateLink)
}