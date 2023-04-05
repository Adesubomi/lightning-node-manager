package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"strings"
)

const (
	Version = "1.0"
)

type AppEnv string

func (e AppEnv) Is(env AppEnv) bool {
	return env == e
}

const AppEnvStaging = "staging"
const AppEnvProduction = "production"   // prod
const AppEnvDevelopment = "development" // develop
const AppEnvTesting = "testing"         // test

type DatabaseConfig struct {
	Connection string
	Host       string
	Port       string
	DbName     string `toml:"db_name"`
	User       string `toml:"user"`
	Password   string
}

type RedisConfig struct {
	Host     string
	Port     string
	DbName   string
	User     string `toml:"user"`
	Password string
}

type SentryConfig struct {
	DNS     string `toml:"dns"`
	Debug   string `toml:"debug"`
	Release string `toml:"release"`
}

type LightningConfig struct {
	Url      string `toml:"url"`
	Port     string `toml:"port"`
	TlsCert  string `toml:"tls_cert"`
	Macaroon string `toml:"macaroon"`
}

type Config struct {
	AppName  string         `toml:"app_name"`
	AppPort  string         `toml:"app_port"`
	AppKey   string         `toml:"app_key"`
	AppENV   string         `toml:"app_env"`
	Database DatabaseConfig `toml:"database"`
	Redis    RedisConfig    `toml:"redis"`
	Sentry   SentryConfig   `toml:"sentry"`
}

// GetEnv get the environment we're running in
// and defaults to use staging environment
func (c *Config) GetEnv() AppEnv {
	switch true {
	case strings.EqualFold(c.AppENV, "development"),
		strings.EqualFold(c.AppENV, "develop"):
		return AppEnvDevelopment
	case strings.EqualFold(c.AppENV, "production"),
		strings.EqualFold(c.AppENV, "prod"):
		return AppEnvProduction
	case strings.EqualFold(c.AppENV, "testing"),
		strings.EqualFold(c.AppENV, "test"):
		return AppEnvTesting
	default:
		return AppEnvStaging
	}
}

func LoadConfigFromFile(filePath string) (*Config, error) {
	config := &Config{}
	if _, err := toml.DecodeFile(filePath, config); err != nil {
		fmt.Println("Error decoding TOML file:", err)
		return nil, err
	}

	return config, nil
}
