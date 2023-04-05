package server

import (
	configPkg "github.com/Adesubomi/lightning-node-manager/pkg/config"
	"github.com/gofiber/fiber/v2"
)

func RegisterBackendRoutes(http *configPkg.HttpStub) *fiber.App {
	app := fiber.New()

	// get list of backend connections
	app.Get("/connections")

	// test a backend connections
	app.Post("/connections/:backend/test")

	// get a list of lightning nodes connected to a backend
	app.Get("/connections/:backend/ln")

	// get the details of a lightning node
	app.Get("/ln/:ln")

	// test lightning node connection/status/availability
	app.Post("/ln/:ln/test")

	// view the details of a lightning node
	app.Get("/ln/:ln/config")

	// create lightning node configuration on a backend
	app.Post("/ln/:ln/config")

	// get list of connected peers to lightning node
	app.Get("/ln/:ln/peers")

	// connect a lightning node to a peer
	app.Post("/ln/:ln/peers")

	// get list of open lightning payment channels
	app.Get("/ln/:ln/channels")

	// open a lightning payment channel
	app.Post("/ln/:ln/channels")

	// close a lightning payment channel
	app.Post("/ln/:ln/channels/close")

	// close a lightning payment channel
	app.Post("/ln/:ln/channels/force-close")

	http.Router.Mount("/backend", app)
	return app
}
