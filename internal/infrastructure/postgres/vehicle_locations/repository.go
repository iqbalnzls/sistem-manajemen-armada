package vehicle_locations

import (
	"context"

	"github.com/iqbalnzls/sistem-manajemen-armada/internal/domain"
)

type Repository interface {
	Insert(ctx context.Context, vehicleLoc *domain.VehicleLocations) error
	FindBy(ctx context.Context, query string, args map[string]interface{}) (*domain.VehicleLocations, error)
}
