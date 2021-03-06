package links

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/mustafasegf/go-shortener/entity"
	"gorm.io/gorm"
)

type Controller struct {
	svc *Service
}

func NewController(svc *Service) *Controller {
	return &Controller{
		svc: svc,
	}
}

func (ctrl *Controller) CreateLink(ctx *gin.Context) {
	req := entity.CreateLinkRequest{}
	err := ctx.ShouldBindBodyWith(&req, binding.JSON)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, entity.Message(err.Error()))
		return
	}
	valid := ctrl.svc.CheckURL(req.LongUrl)
	if !valid {
		ctx.IndentedJSON(http.StatusBadRequest, entity.Message("Not valid URL"))
		return
	}

	_, err = ctrl.svc.GetLinkByURL(req.ShortUrl)
	if err == nil {
		ctx.IndentedJSON(http.StatusConflict, entity.Message("Short Url Exist"))
		return
	} else if err != gorm.ErrRecordNotFound {
		ctx.IndentedJSON(http.StatusInternalServerError, entity.Message(err.Error()))
		log.Print(err.Error())
		return
	}

	err = ctrl.svc.InsertURL(req)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, entity.Message(err.Error()))
		log.Print(err.Error())
		return
	}

	ctx.IndentedJSON(http.StatusOK, entity.Message("Success"))
	return
}

func (ctrl *Controller) Redirect(ctx *gin.Context) {
	shortUrl := ctx.Param("url")
	if shortUrl == "" {
		ctx.Redirect(http.StatusNotFound, "/")
		return
	}

	data, err := ctrl.svc.GetLinkByURL(shortUrl)
	if err == gorm.ErrRecordNotFound {
		ctx.Redirect(http.StatusTemporaryRedirect, "/")
		return
	} else if err != nil {
		ctx.Redirect(http.StatusTemporaryRedirect, "/")
		log.Print(err.Error())
		return
	}

	ctx.Redirect(http.StatusFound, data.LongUrl)
	return
}
