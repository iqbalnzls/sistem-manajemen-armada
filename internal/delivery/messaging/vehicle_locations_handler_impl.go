package messaging

import (
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/iqbalnzls/sistem-manajemen-armada/internal/common/logger"
	"github.com/iqbalnzls/sistem-manajemen-armada/internal/common/validator"
	"github.com/iqbalnzls/sistem-manajemen-armada/internal/dto"
	"github.com/iqbalnzls/sistem-manajemen-armada/internal/usecase/vehiclelocations"
)

type vehicleLocHandler struct {
	svc    vehiclelocations.Service
	logger *zap.Logger
	v      *validator.Validator
}

func NewVehicleLocationHandler(svc vehiclelocations.Service, log *zap.Logger, v *validator.Validator) VehicleLocationHandler {
	if svc == nil {
		panic("svc is nil")
	}
	if log == nil {
		panic("logger is nil")
	}
	if v == nil {
		panic("validator is nil")
	}

	return &vehicleLocHandler{
		svc:    svc,
		logger: log,
		v:      v,
	}
}

func (h *vehicleLocHandler) ReceiveVehicleLocation(topic string, payload []byte) error {
	tn := time.Now()
	log := h.logger
	log = log.With(
		zap.String("xid", uuid.New().String()),
	)

	ctx := logger.WithLogger(context.Background(), log)

	log.Info("Incoming messaging",
		zap.String("tag", "T1"),
		zap.String("timestamp", time.Now().Format(time.RFC3339)),
		zap.Any("payload", string(payload)),
	)

	var vehicleLocation *dto.ReceiveVehicleLocation

	if err := json.Unmarshal(payload, &vehicleLocation); err != nil {
		log.Error("Failed to unmarshal vehicle location payload", zap.Error(err))
		return err
	}

	if err := h.v.Validate(vehicleLocation); err != nil {
		log.Error("Invalid vehicle location payload", zap.Error(err))
		return err
	}

	if err := h.svc.ReceiveVehicleLocation(ctx, vehicleLocation); err != nil {
		return err
	}

	log.Info("Outgoing messaging",
		zap.String("tag", "T4"),
		zap.Int64("rt", time.Now().Sub(tn).Milliseconds()))

	return nil
}
