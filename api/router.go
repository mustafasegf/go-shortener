package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mustafasegf/go-shortener/controller"
	"github.com/mustafasegf/go-shortener/repository"
	"github.com/mustafasegf/go-shortener/service"
)

type Route struct {
	router *gin.Engine
}

func (s *Server) setupRouter() {

	linkRepo := repository.NewLinkRepo(s.db)
	linkSvc := service.NewLinkService(linkRepo)
	linkCtlr := controller.NewLinkController(linkSvc)

	s.router.GET("/:url", linkCtlr.Redirect)

	api := s.router.Group("/api")
	link := api.Group("/link")

	link.POST("create", linkCtlr.CreateLink)
	
}
