package router

import (
	//"encoding/json"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pester18/url-shortener/interface/controller"
)

func NewRouter(r *gin.Engine, c controller.AppController) *gin.Engine {
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/:url_token", func(context *gin.Context) {
		url_token := context.Param("url_token")

		c.RedirectToFullUrl(url_token, context)
	})

	r.POST("/shorten", func(context *gin.Context) {
		var reqBody struct {
			Url string
		}

		err := context.BindJSON(&reqBody)
		if err != nil {
			context.String(
				http.StatusBadRequest,
				err.Error(),
			)
		}
		c.GenerateShortUrl(reqBody.Url, context)
	})

	return r
}
