package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"github.com/pester18/url-shortener/config"
	"github.com/pester18/url-shortener/infrastructure/datastore"
	"github.com/pester18/url-shortener/infrastructure/router"
	"github.com/pester18/url-shortener/registry"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func main() {
	conf := config.New()

	dialStr := fmt.Sprintf("%s:%s", conf.Mongo.Host, conf.Mongo.Port)
	db := datastore.NewDB(dialStr, conf.Mongo.DB)

	reg := registry.NewRegistry(db)

	r := router.NewRouter(reg.NewAppController())
	srv := &http.Server{
		Addr:    ":" + conf.Server.Port,
		Handler: r,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}
