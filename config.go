package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

type CrawlerInfo struct {
	URL string `toml:"url"`
}

func Configuration() (config *CrawlerInfo) {
	// Read config
	// All config stored in config.toml, and here is an example how to
	// decode config item from toml format to type that we have defined
	// in struct
	config := CrawlerInfo{}
	_, err := toml.DecodeFile("./config.toml", &config)
	if err != nil {
		log.Fatalf(ErrConfDecode, err)
	}

	return &config
}
