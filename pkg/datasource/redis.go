package datasource

import (
	"fmt"
	configPkg "github.com/Adesubomi/lightning-node-manager/pkg/config"
	logPkg "github.com/Adesubomi/lightning-node-manager/pkg/log"
	"github.com/go-redis/redis"
	"log"
)

func connectByString(redisConf *configPkg.RedisConfig) (*redis.Client, error) {
	options, err := redis.ParseURL(fmt.Sprintf(
		"rediss://%v:%v@%v:%v",
		redisConf.User,
		redisConf.Password,
		redisConf.Host,
		redisConf.Port,
	))

	if err != nil {
		return nil, err
	}

	return redis.NewClient(options), nil
}

func connectByOptions(redisConf *configPkg.RedisConfig) (*redis.Client, error) {
	return redis.NewClient(&redis.Options{
		Network: "",
		Addr: fmt.Sprintf(
			"%v:%v", // localhost:6379
			redisConf.Host,
			redisConf.Port,
		),
		Password: redisConf.Password,
		DB:       0,
	}), nil
}

func RedisConnection(redisConf *configPkg.RedisConfig) *redis.Client {
	var client *redis.Client
	var err error

	if redisConf.User != "" {
		client, err = connectByString(redisConf)
	} else {
		client, err = connectByOptions(redisConf)
	}

	if err != nil {
		msg := fmt.Sprintf(" ?? Could not connect to Redis: %v\n", err)
		logPkg.PrintlnRed(msg)
		log.Fatal(err)
		return nil
	}

	if _, err = client.Ping().Result(); err != nil {
		msg := fmt.Sprintf(" ?? Could not connect to Redis because: %v\n", err)
		logPkg.PrintlnRed(msg)
		log.Fatal(err)
		return nil
	}

	logPkg.PrintlnGreen("  ✔ Redis Connection Established")
	return client
}
