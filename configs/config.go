package configs

import (
	"github.com/BurntSushi/toml"
	"log"
	"os"
)

type Config struct {
	Application struct {
		Name string `toml:"name"`
	}

	Server struct {
		Port string `toml:"port"`
	}
}

func NewConfig(path string) Config {
	configRaw, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("reading config %s error: %v", path, err)
	}

	config := Config{}
	err = toml.Unmarshal(configRaw, &config)
	if err != nil {
		log.Fatalf("parsing config %s error: %v", path, err)
	}

	return config
}
