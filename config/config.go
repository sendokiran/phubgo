package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

type Configuration struct {
	CrawlerInfo CrawlerInfo `toml:"crawler"`
}

type CrawlerInfo struct {
	URL string `toml:"url"`
}

const (
	ConfigPath    string = "./config.toml"
	ErrConfDecode string = "Failed to decode config: %v"
)

var (
	ConfigurationData = &Configuration{}
)

func init() {
	// Read config
	// All config stored in config.toml, and here is an example how to
	// decode config item from toml format to type that we have defined
	// in struct
	_, err := toml.DecodeFile(ConfigPath, ConfigurationData)
	if err != nil {
		log.Fatalf(ErrConfDecode, err)
	}
}

func New() *Configuration {
	return ConfigurationData
}
