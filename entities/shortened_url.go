package entities

import "github.com/lithammer/shortuuid"

type ShortenedURL struct {
	URL      string `json:"url" bson:"url"`
	URLtoken string `json:"url_token" bson:"url_token"`
}

func CreateShortenedUrl(url string) *ShortenedURL {
	token := shortuuid.New()
	return &ShortenedURL{
		URL: url,
		URLtoken: token,
	}
}

func (*ShortenedURL) CollectionName() string { return "urls" }
