package entities

//import "gopkg.in/mgo.v2/bson"

type ShortenedURL struct {
	URL      string `json:"url" bson:"url"`
	URLtoken string `json:"url_token" bson:"url_token"`
}

func (*ShortenedURL) CollectionName() string { return "urls" }
