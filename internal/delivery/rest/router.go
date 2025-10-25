package rest

import "github.com/gofiber/fiber/v2"

func SetupRouter(app *fiber.App, handler *Handler) {
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("service is up and running...")
	})

	vehicle := app.Group("/vehicles")
	vehicle.Get("/:vehicle_id/location", handler.vehicleLoc.FindVehicleById)
	vehicle.Get("/:vehicle_id/history", handler.vehicleLoc.FindVehicleByIdAndTime)

}
