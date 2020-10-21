package controller

import (
	"github.com/pester18/url-shortener/entities"
	"github.com/pester18/url-shortener/usecase/interactor"
)

type appController struct {
	interact interactor.Interactor
}

type AppController interface {
	GenerateShortUrl(url string) Response
	GetFullUrl(urlToken string) Response
	DeleteShortUrl(urlToken string) Response
}

type Response struct {
	Success bool        `json:"success"`
	Error   string      `json:"error,omitempty"`
	Result  interface{} `json:"result,omitempty"`
}

func NewAppController(interact interactor.Interactor) AppController {
	return &appController{interact}
}

func (ac *appController) GenerateShortUrl(url string) Response {
	shortenedUrlToCreate := &entities.ShortenedURL{URL: url}

	shortenedUrl, err := ac.interact.CreateShortenedUrl(shortenedUrlToCreate)

	if err != nil {
		return Response{
			Success: false,
			Error:   err.Error(),
		}
	}

	return Response{
		Success: true,
		Result:  shortenedUrl,
	}
}

func (ac *appController) GetFullUrl(urlToken string) Response {
	shortenedUrlToFind := &entities.ShortenedURL{URLtoken: urlToken}

	shortenedUrl, err := ac.interact.GetShortenedUrlOrigin(shortenedUrlToFind)

	if err != nil {
		return Response{
			Success: false,
			Error:   err.Error(),
		}
	}

	return Response{
		Success: true,
		Result:  shortenedUrl.URL,
	}
}

func (ac *appController) DeleteShortUrl(urlToken string) Response {
	shortenedUrlToDelete := &entities.ShortenedURL{URLtoken: urlToken}

	err := ac.interact.RemoveShortenedUrl(shortenedUrlToDelete)

	if err != nil {
		return Response{
			Success: false,
			Error:   err.Error(),
		}
	}

	return Response{
		Success: true,
	}
}
