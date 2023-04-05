package main

import (
	"fmt"
	configPkg "github.com/Adesubomi/magic-ayo-api/pkg/config"
	logPkg "github.com/Adesubomi/magic-ayo-api/pkg/log"
	networkPkg "github.com/Adesubomi/magic-ayo-api/pkg/network"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {

	conf := configPkg.LoadConfig("cmd/config.toml")

	// sentry connection setup
	_ = logPkg.ConnectToSentry(conf.Sentry, conf.GetEnv())

	app := fiber.New()
	app.Use(logPkg.FiberRequestDebug)
	app.Use(networkPkg.CorsFiberMiddleware)

	err := app.Listen(fmt.Sprintf(":%v", conf.AppPort))
	if err != nil {
		log.Fatalf("server listen failed, %s", err.Error())
	}
}
