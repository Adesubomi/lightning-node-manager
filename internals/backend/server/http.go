package server

import (
	configPkg "github.com/Adesubomi/lightning-node-manager/pkg/config"
	"github.com/gofiber/fiber/v2"
)

func RegisterBackendRoutes(*configPkg.HttpStub) *fiber.App {
	app := fiber.New()

	return app
}
