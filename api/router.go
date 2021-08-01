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

	s.router.LoadHTMLGlob("templates/*")
	s.router.Static("/static", "./static")

	linkRepo := repository.NewLinkRepo(s.db, s.rdb)
	linkSvc := service.NewLinkService(linkRepo)
	linkCtlr := controller.NewLinkController(linkSvc)

	staticCtlr := controller.NewStaticController()

	s.router.GET("/:url", linkCtlr.Redirect)
	s.router.GET("/", staticCtlr.Index)

	api := s.router.Group("/api")
	link := api.Group("/link")

	link.POST("create", linkCtlr.CreateLink)

}
