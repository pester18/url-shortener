package repository

import (
	"fmt"

	"github.com/pester18/url-shortener/entities"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type repository struct {
	db *mgo.Database
}

type Repository interface {
	FindShortenedUrl(shortenedUrl *entities.ShortenedURL) (*entities.ShortenedURL, error)
	SaveShortenedUrl(shortenedUrl *entities.ShortenedURL) error
	DeleteShortenedUrl(shortenedUrl *entities.ShortenedURL) error
}

func NewRepository(db *mgo.Database) Repository {
	return &repository{db}
}

func (r *repository) FindShortenedUrl(shortenedUrl *entities.ShortenedURL) (*entities.ShortenedURL, error) {
	if shortenedUrl == nil {
		return nil, fmt.Errorf("Error: no shortened url provided")
	}

	collection := r.db.C(shortenedUrl.CollectionName())

	res := entities.ShortenedURL{}

	err := collection.Find(bson.M{"url_token": shortenedUrl.URLtoken}).One(&res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *repository) SaveShortenedUrl(shortenedUrl *entities.ShortenedURL) error {
	if shortenedUrl == nil {
		return fmt.Errorf("Error: no shortened url provided")
	}

	collection := r.db.C(shortenedUrl.CollectionName())

	err := collection.Insert(shortenedUrl)

	return err
}

func (r *repository) DeleteShortenedUrl(shortenedUrl *entities.ShortenedURL) error {
	if shortenedUrl == nil {
		return fmt.Errorf("Error: no shortened url provided")
	}

	collection := r.db.C(shortenedUrl.CollectionName())

	err := collection.Remove(bson.M{"url_token": shortenedUrl.URLtoken})

	return err
}
