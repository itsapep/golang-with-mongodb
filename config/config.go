package config

import "os"

type APIConfig struct {
	APIPort string
	APIHost string
}

type MongoConfig struct {
	Host     string
	Port     string
	DbName   string
	User     string
	Password string
}

type Config struct {
	APIConfig
	MongoConfig
}

func (c *Config) readConfig() {
	c.MongoConfig = MongoConfig{
		Host:     os.Getenv("MONGO_HOST"),     //127.0.0.1
		Port:     os.Getenv("MONGO_PORT"),     //27017
		DbName:   os.Getenv("MONGO_DB"),       //contohDb
		User:     os.Getenv("MONGO_USER"),     //yurhamafif
		Password: os.Getenv("MONGO_PASSWORD"), //12345678
	}
	c.APIConfig = APIConfig{
		APIHost: os.Getenv("API_HOST"), //localhost
		APIPort: os.Getenv("API_PORT"), //8888
	}
}

func NewConfig() Config {
	cfg := Config{}
	cfg.readConfig()
	return cfg
}
