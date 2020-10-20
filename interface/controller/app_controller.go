package controller

import (
	"net/http"

	"github.com/pester18/url-shortener/entities"
	"github.com/pester18/url-shortener/usecase/interactor"
)

type appController struct {
	interact interactor.Interactor
}

type AppController interface {
	GenerateShortUrl(url string, c Context)
	RedirectToFullUrl(url_token string, c Context)
}

type failureResponse struct {
	Error string
}

func NewAppController(interact interactor.Interactor) AppController {
	return &appController{interact}
}

func (ac *appController) GenerateShortUrl(url string, c Context) {
	shortenedUrlToCreate := &entities.ShortenedURL{URL: url}

	res, err := ac.interact.CreateShortenedUrl(shortenedUrlToCreate)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			failureResponse{
				Error: err.Error(),
			},
		)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (ac *appController) RedirectToFullUrl(url_token string, c Context) {
	shortenedUrlToFind := &entities.ShortenedURL{URLtoken: url_token}

	res, err := ac.interact.GetShortenedUrl(shortenedUrlToFind)

	if err != nil {
		c.JSON(
			http.StatusNotFound,
			failureResponse{
				Error: err.Error(),
			},
		)
		return
	}

	c.Redirect(http.StatusFound, res.URL)
}
