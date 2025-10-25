package messaging

import (
	"go.uber.org/zap"

	"github.com/iqbalnzls/sistem-manajemen-armada/internal/container"
)

type Handler struct {
	vehicleLoc VehicleLocationHandler
}

func NewMessagingHandler(c *container.Container, logger *zap.Logger) *Handler {
	return &Handler{
		vehicleLoc: NewVehicleLocationHandler(c.VehicleLocSvc, logger, c.Validator),
	}
}
