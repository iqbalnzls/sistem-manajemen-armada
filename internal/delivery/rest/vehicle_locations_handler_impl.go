package rest

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"github.com/iqbalnzls/sistem-manajemen-armada/internal/common/logger"
	"github.com/iqbalnzls/sistem-manajemen-armada/internal/common/validator"
	"github.com/iqbalnzls/sistem-manajemen-armada/internal/dto"
	"github.com/iqbalnzls/sistem-manajemen-armada/internal/usecase/vehiclelocations"
)

type vehicleLocHandler struct {
	svc vehiclelocations.Service
	v   *validator.Validator
}

func NewVehicleLocationsHandler(svc vehiclelocations.Service, v *validator.Validator) VehicleLocationHandler {
	if svc == nil {
		panic("svc is nil")
	}
	if v == nil {
		panic("validator is nil")
	}

	return &vehicleLocHandler{
		svc: svc,
		v:   v,
	}
}

func (h *vehicleLocHandler) FindVehicleById(c *fiber.Ctx) (err error) {
	log := logger.FromContext(c.UserContext())
	ctx := c.UserContext()

	req := &dto.FindVehicleByIdRequest{
		VehicleId: c.Params("vehicle_id"),
	}

	if err = h.v.Validate(req); err != nil {
		log.Error("Invalid request", zap.Error(err))
		return
	}

	resp, err := h.svc.FindVehicleById(ctx, req)
	if err != nil {
		return
	}

	startTime := ctx.Value("startTime").(time.Time)
	log.Info("Outgoing request",
		zap.String("tag", "T4"),
		zap.Int64("rt", time.Now().Sub(startTime).Milliseconds()),
	)

	return c.Status(fiber.StatusOK).JSON(resp)
}

func (h *vehicleLocHandler) FindVehicleByIdAndTime(c *fiber.Ctx) (err error) {
	log := logger.FromContext(c.UserContext())
	ctx := c.UserContext()

	req := &dto.FindVehicleByIdAndTimeRequest{
		VehicleId: c.Params("vehicle_id"),
	}

	if err = c.QueryParser(req); err != nil {
		log.Error("Failed to parse query params", zap.Error(err))
		return
	}

	if err = h.v.Validate(req); err != nil {
		log.Error("Invalid request", zap.Error(err))
		return
	}

	resp, err := h.svc.FindVehicleByIdAndTime(ctx, req)
	if err != nil {
		return
	}

	startTime := ctx.Value("startTime").(time.Time)
	log.Info("Outgoing request",
		zap.String("tag", "T4"),
		zap.Int64("rt", time.Now().Sub(startTime).Milliseconds()),
	)

	return c.Status(fiber.StatusOK).JSON(resp)
}
