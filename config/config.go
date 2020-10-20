package config

import (
	"os"
)

// Config contains the configuration of the url shortener.
type ServerConfig struct {
	Port string
}

type MongoConfig struct {
	Host string
	Port string
	DB   string
}

type Config struct {
	Server ServerConfig
	Mongo  MongoConfig
}

func New() *Config {
	return &Config{
		Server: ServerConfig{
			Port: getEnv("PORT", "8080"),
		},
		Mongo: MongoConfig{
			Host: getEnv("MONGO_HOST", "localhost"),
			Port: getEnv("MONGO_PORT", "27017"),
			DB:   getEnv("MONGO_DB_NAME", "url_shortener"),
		},
	}
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
