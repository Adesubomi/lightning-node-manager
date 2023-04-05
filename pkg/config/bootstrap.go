package config

import (
	"flag"
	"fmt"
	"log"
)

func LoadConfig(filePath string) *Config {
	configFile := flag.String(
		"config", filePath,
		"path to config.toml config file")

	flag.Parse()
	conf, err := LoadConfigFromFile(*configFile)
	if err != nil {
		fmt.Println("")
		log.Fatalf("\n    ✗ failed to load config file\n    error: %v", err.Error())
	}

	return conf
}
