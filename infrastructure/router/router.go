package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pester18/url-shortener/interface/controller"
)

type reqBody struct {
	Url string `json:"url" binding:"required"`
}

func NewRouter(c controller.AppController) http.Handler {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/:url_token", func(context *gin.Context) {
		urlToken := context.Param("url_token")

		res := c.GetFullUrl(urlToken)

		if !res.Success {
			context.JSON(
				http.StatusNotFound,
				res,
			)
			return
		}

		redirectUrl, ok := res.Result.(string)
		if ok {
			context.Redirect(http.StatusFound, redirectUrl)
			return
		}

		context.JSON(
			http.StatusInternalServerError,
			controller.Response{
				Success: false,
			},
		)
	})

	r.DELETE("/:url_token", func(context *gin.Context) {
		urlToken := context.Param("url_token")

		res := c.DeleteShortUrl(urlToken)

		if !res.Success {
			context.JSON(
				http.StatusNotFound,
				res,
			)
			return
		}

		context.JSON(
			http.StatusOK,
			res,
		)
	})

	r.POST("/shorten", func(context *gin.Context) {
		req := reqBody{}

		err := context.BindJSON(&req)
		if err != nil {
			context.JSON(
				http.StatusBadRequest,
				controller.Response{
					Success: false,
					Error:   err.Error(),
				},
			)
			return
		}

		res := c.GenerateShortUrl(req.Url)

		if !res.Success {
			context.JSON(
				http.StatusBadRequest,
				res,
			)
			return
		}

		context.JSON(
			http.StatusOK,
			res,
		)
	})

	return r
}
