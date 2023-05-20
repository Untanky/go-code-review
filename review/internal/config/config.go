package config

import (
	"log"

	"github.com/brumhard/alligotor"
)

type ApiConfig struct {
	Host string
	Port int
}

type Config struct {
	API ApiConfig
}

func New() Config {
	cfg := Config{}
	if err := alligotor.Get(&cfg); err != nil {
		log.Fatal(err)
	}
	return cfg
}
