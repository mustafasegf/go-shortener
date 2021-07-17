package api

import (
	"github.com/mustafasegf/go-shortener/controller"
	"github.com/gin-gonic/gin"
)
type Route struct {
	router *gin.Engine
}


func (s *Server) setupRouter() {
	api := s.router.Group("/api")

	link := api.Group("/link")

	link.POST("create", controller.CreateLink)
}