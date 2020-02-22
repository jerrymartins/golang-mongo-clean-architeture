package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

type Config struct {
	UrlApi                 string
	GoogleClientId         string
	GoogleClientSecret     string
	GoogleApiTokenValidate string
	MongoUri               string
	Database               string
	SecretKey              string
}

func (c *Config) Read() {
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatal(err)
	}
}
