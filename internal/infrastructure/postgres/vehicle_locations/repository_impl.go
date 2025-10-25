package vehicle_locations

import (
	"context"

	"go.uber.org/zap"

	"github.com/iqbalnzls/sistem-manajemen-armada/internal/domain"
	"github.com/iqbalnzls/sistem-manajemen-armada/pkg/database"
	"github.com/iqbalnzls/sistem-manajemen-armada/pkg/logger"
)

type repo struct {
	db *database.Database
}

func NewVehicleLocationsRepository(db *database.Database) Repository {
	if db == nil {
		panic("db is nil")
	}

	return &repo{db: db}
}

func (r *repo) Insert(ctx context.Context, vehicleLoc *domain.VehicleLocations) error {
	err := r.db.Create(vehicleLoc).Error
	if err != nil {
		logger.FromContext(ctx).Error("failed to save vehicle location", zap.Error(err))
		return err
	}

	return nil
}

func (r *repo) FindBy(ctx context.Context, query string, args map[string]interface{}) (*domain.VehicleLocations, error) {
	var result domain.VehicleLocations

	err := r.db.Where(query, args).Order("timestamp DESC").First(&result).Error
	if err != nil {
		logger.FromContext(ctx).Error("failed to find vehicle location", zap.Error(err))
	}

	return &result, err
}
