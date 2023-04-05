package config

import (
	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type HttpStub struct {
	Conf        *Config
	RedisClient *redis.Client
	DbClient    *gorm.DB
	Router      *fiber.App
}
