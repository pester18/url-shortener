package interactor

import (
	"fmt"
	"github.com/pester18/url-shortener/entities"
	"github.com/pester18/url-shortener/usecase/repository"
	"github.com/pester18/url-shortener/utils"
)

type interactor struct {
	Repository repository.Repository
}

type Interactor interface {
	GetShortenedUrl(shortenedUrlToFind *entities.ShortenedURL) (*entities.ShortenedURL, error)
	CreateShortenedUrl(shortenedUrlToCreate *entities.ShortenedURL) (*entities.ShortenedURL, error)
}

func NewInteractor(r repository.Repository) Interactor {
	return &interactor{r}
}

func (i *interactor) GetShortenedUrl(shortenedUrlToFind *entities.ShortenedURL) (*entities.ShortenedURL, error) {
	res, err := i.Repository.FindShortenedUrl(shortenedUrlToFind)
	if err != nil {
		return nil, err
	}

	if res.URL == "" {
		return nil, fmt.Errorf("There is no saved url with token: %s", shortenedUrlToFind.URLtoken)
	}

	return res, nil
}

func (i *interactor) CreateShortenedUrl(shortenedUrlToCreate *entities.ShortenedURL) (*entities.ShortenedURL, error) {
	urlToken := utils.GenerateToken()

	shortenedUrlToCreate.URLtoken = urlToken

	err := i.Repository.SaveShortenedUrl(shortenedUrlToCreate)
	if err != nil {
		return nil, err
	}

	return shortenedUrlToCreate, nil
}
