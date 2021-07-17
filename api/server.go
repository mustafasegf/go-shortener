package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mustafasegf/go-shortener/util"
	"gorm.io/gorm"
)

type Server struct {
	config util.Config
	router *gin.Engine
	db     *gorm.DB
}

func MakeServer(config util.Config, db *gorm.DB) Server {
	router := gin.Default()
	server := Server{
		config: config,
		router: router,
		db:     db,
	}
	return server
}

func (s *Server) RunServer() {
	s.setupRouter()
	serverString := fmt.Sprintf(":%s", s.config.ServerPort)
	s.router.Run(serverString)
}
