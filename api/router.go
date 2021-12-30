package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mustafasegf/go-shortener/links"
)

type Route struct {
	router *gin.Engine
}

func (s *Server) setupRouter() {
	s.router.LoadHTMLGlob("templates/*")
	s.router.Static("/static", "./static")

	linkRepo := links.NewRepo(s.db, s.rdb)
	linkSvc := links.NewService(linkRepo)
	linkCtlr := links.NewController(linkSvc)

	s.router.GET("/:url", linkCtlr.Redirect)
	s.router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", "")
	})

	api := s.router.Group("/api")
	link := api.Group("/link")

	link.POST("create", linkCtlr.CreateLink)

}
