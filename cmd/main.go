package main

import (
	"fmt"
	backendHttp "github.com/Adesubomi/lightning-node-manager/internals/backend/server"
	configPkg "github.com/Adesubomi/lightning-node-manager/pkg/config"
	dataPkg "github.com/Adesubomi/lightning-node-manager/pkg/datasource"
	logPkg "github.com/Adesubomi/lightning-node-manager/pkg/log"
	networkPkg "github.com/Adesubomi/lightning-node-manager/pkg/network"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {

	conf := configPkg.LoadConfig("cmd/config.toml")

	// sentry connection setup
	_ = logPkg.ConnectToSentry(conf.Sentry, conf.GetEnv())

	dbClient := dataPkg.ConnectDatabase(&conf.Database)

	redisClient := dataPkg.RedisConnection(&conf.Redis)

	app := fiber.New()
	app.Use(logPkg.FiberRequestDebug)
	app.Use(networkPkg.CorsFiberMiddleware)

	httpStub := &configPkg.HttpStub{
		Conf:        conf,
		DbClient:    dbClient,
		RedisClient: redisClient,
		Router:      app,
	}

	// Backend
	backendHttp.RegisterBackendRoutes(httpStub)

	err := app.Listen(fmt.Sprintf(":%v", conf.AppPort))
	if err != nil {
		log.Fatalf("server listen failed, %s", err.Error())
	}
}
