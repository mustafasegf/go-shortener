package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/mustafasegf/go-shortener/entity"
	"github.com/mustafasegf/go-shortener/service"
)

type Link struct {
	svc *service.Link
}

func NewLinkController(svc *service.Link) *Link {
	return &Link{
		svc: svc,
	}
}

func (ctrl *Link) CreateLink(ctx *gin.Context) {
	req := entity.CreateLinkRequest{}
	err := ctx.ShouldBindBodyWith(&req, binding.JSON)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, entity.Message(err.Error()))
		return
	}
	ctrl.svc
}
