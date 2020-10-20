package repository

import "github.com/pester18/url-shortener/entities"

type Repository interface {
	FindShortenedUrl(shortenedUrl *entities.ShortenedURL) (*entities.ShortenedURL, error)
	SaveShortenedUrl(shortenedUrl *entities.ShortenedURL) error
	// FindAllUrls() ([]*entities.ShortenedURL, error)
}
