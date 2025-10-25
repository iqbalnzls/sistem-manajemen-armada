package rest

import (
	"github.com/iqbalnzls/sistem-manajemen-armada/internal/container"
)

type Handler struct {
	vehicleLoc VehicleLocationHandler
}

func NewRestHandler(c *container.Container) *Handler {
	return &Handler{
		vehicleLoc: NewVehicleLocationsHandler(c.VehicleLocSvc, c.Validator),
	}

}
