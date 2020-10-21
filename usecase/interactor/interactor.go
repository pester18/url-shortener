package interactor

import (
	"fmt"

	"github.com/pester18/url-shortener/entities"
	"github.com/pester18/url-shortener/usecase/repository"
)

type interactor struct {
	Repository repository.Repository
}

type Interactor interface {
	GetShortenedUrlOrigin(shortenedUrlToFind *entities.ShortenedURL) (*entities.ShortenedURL, error)
	CreateShortenedUrl(shortenedUrlToCreate *entities.ShortenedURL) (*entities.ShortenedURL, error)
	RemoveShortenedUrl(shortenedUrlToDelete *entities.ShortenedURL) error
}

func NewInteractor(r repository.Repository) Interactor {
	return &interactor{r}
}

func (i *interactor) GetShortenedUrlOrigin(shortenedUrlToFind *entities.ShortenedURL) (*entities.ShortenedURL, error) {
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
	shortenedUrl := entities.CreateShortenedUrl(shortenedUrlToCreate.URL)

	err := i.Repository.SaveShortenedUrl(shortenedUrl)
	if err != nil {
		return nil, err
	}

	return shortenedUrl, nil
}

func (i *interactor) RemoveShortenedUrl(shortenedUrlToDelete *entities.ShortenedURL) error {
	if shortenedUrlToDelete.URLtoken == "" {
		return fmt.Errorf("No url token provided")
	}

	err := i.Repository.DeleteShortenedUrl(shortenedUrlToDelete)

	return err
}
