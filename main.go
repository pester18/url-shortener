package main

import (
	//"encoding/json"

	"fmt"
	"log"
	//"net/http"

	"github.com/gin-gonic/gin"
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

	r := gin.New()
	r = router.NewRouter(r, reg.NewAppController())

	if err := r.Run(":" + conf.Server.Port); err != nil {
		log.Fatalln(err)
	}
}
