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
	RedirectToFullUrl(urlToken string, c Context)
	DeleteShortUrl(urlToken string, c Context)
}

type response struct {
	Success bool `json:"success"`
	Error string `json:"error,omitempty"`
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
			response{
				Success: false,
				Error: err.Error(),
			},
		)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (ac *appController) RedirectToFullUrl(urlToken string, c Context) {
	shortenedUrlToFind := &entities.ShortenedURL{URLtoken: urlToken}

	res, err := ac.interact.GetShortenedUrlOrigin(shortenedUrlToFind)

	if err != nil {
		c.JSON(
			http.StatusNotFound,
			response{
				Success: false,
				Error: err.Error(),
			},
		)
		return
	}

	c.Redirect(http.StatusFound, res.URL)
}

func (ac *appController) DeleteShortUrl(urlToken string, c Context) {
	shortenedUrlToDelete := &entities.ShortenedURL{URLtoken: urlToken}

	err := ac.interact.RemoveShortenedUrl(shortenedUrlToDelete)

	if err != nil {
		c.JSON(
			http.StatusNotFound,
			response{
				Success: false,
				Error: err.Error(),
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		response{
			Success: true,
		},
	)
}
