package vehiclelocations

import (
	"context"

	"github.com/iqbalnzls/sistem-manajemen-armada/internal/dto"
)

type Service interface {
	ReceiveVehicleLocation(ctx context.Context, req *dto.ReceiveVehicleLocation) (err error)
	FindVehicleById(ctx context.Context, req *dto.FindVehicleByIdRequest) (resp dto.FindVehicleByIdResponse, err error)
	FindVehicleByIdAndTime(ctx context.Context, req *dto.FindVehicleByIdAndTimeRequest) (resp dto.FindVehicleByIdAndTimeResponse, err error)
}
