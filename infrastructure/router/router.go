package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pester18/url-shortener/interface/controller"
)

type reqBody struct {
	Url string `json:"url" binding:"required"`
}

func NewRouter(r *gin.Engine, c controller.AppController) *gin.Engine {
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/:url_token", func(context *gin.Context) {
		urlToken := context.Param("url_token")

		c.RedirectToFullUrl(urlToken, context)
	})

	r.DELETE("/:url_token", func(context *gin.Context) {
		urlToken := context.Param("url_token")

		c.DeleteShortUrl(urlToken, context)
	})

	r.POST("/shorten", func(context *gin.Context) {
		req := reqBody{}

		err := context.BindJSON(&req)
		if err != nil {
			context.String(
				http.StatusBadRequest,
				err.Error(),
			)
			return
		}

		c.GenerateShortUrl(req.Url, context)
	})

	return r
}
