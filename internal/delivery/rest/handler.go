package rest

import (
	"github.com/gofiber/fiber/v2"
)

type VehicleLocationHandler interface {
	FindVehicleById(c *fiber.Ctx) (err error)
	FindVehicleByIdAndTime(c *fiber.Ctx) (err error)
}
