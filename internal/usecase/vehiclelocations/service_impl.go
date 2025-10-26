package vehiclelocations

import (
	"context"
	"encoding/json"

	"github.com/iqbalnzls/sistem-manajemen-armada/internal/common/constants"
	"github.com/iqbalnzls/sistem-manajemen-armada/internal/common/util"
	"github.com/iqbalnzls/sistem-manajemen-armada/internal/dto"
	"github.com/iqbalnzls/sistem-manajemen-armada/internal/infrastructure/messaging/rabbitmq"
	vehicleLocationsRepo "github.com/iqbalnzls/sistem-manajemen-armada/internal/infrastructure/postgres/vehicle_locations"
)

type service struct {
	vehicleLocRepo vehicleLocationsRepo.Repository
	rabbitmqPub    rabbitmq.Publisher
}

func NewVehicleLocationsService(vehicleLocRepo vehicleLocationsRepo.Repository, rabbitmqPub rabbitmq.Publisher) Service {
	if vehicleLocRepo == nil {
	}
	if rabbitmqPub == nil {
		panic("rabbitmqPub is nil")
	}

	return &service{
		vehicleLocRepo: vehicleLocRepo,
		rabbitmqPub:    rabbitmqPub,
	}
}

func (s *service) ReceiveVehicleLocation(ctx context.Context, receiveVehicleLoc *dto.ReceiveVehicleLocation) (err error) {
	if err = s.vehicleLocRepo.Insert(ctx, toVehicleLocations(receiveVehicleLoc)); err != nil {
		return
	}

	distance := util.HaversineDistance(constants.JakartaLatitude, constants.JakartaLongitude, receiveVehicleLoc.Latitude, receiveVehicleLoc.Longitude)
	if distance <= constants.GeofenceRadius {
		message := toGeofenceEvent(receiveVehicleLoc)
		b, _ := json.Marshal(message)
		err = s.rabbitmqPub.PublishEvent(ctx, b, rabbitmq.EventTypeGeofenceEntry.String())
	}

	return
}

func (s *service) FindVehicleById(ctx context.Context, req *dto.FindVehicleByIdRequest) (resp dto.FindVehicleByIdResponse, err error) {
	query := "vehicle_id = @vehicle_id"
	args := map[string]interface{}{
		"vehicle_id": req.VehicleId,
	}
	domain, err := s.vehicleLocRepo.FindBy(ctx, query, args)
	if err != nil {
		return
	}

	resp = toFindVehicleByIdResponse(domain)

	return

}

func (s *service) FindVehicleByIdAndTime(ctx context.Context, req *dto.FindVehicleByIdAndTimeRequest) (resp dto.FindVehicleByIdAndTimeResponse, err error) {
	query := "vehicle_id = @vehicle_id AND timestamp >= @start AND timestamp <= @end"
	args := map[string]interface{}{
		"vehicle_id": req.VehicleId,
		"start":      req.Start,
		"end":        req.End,
	}
	domain, err := s.vehicleLocRepo.FindBy(ctx, query, args)
	if err != nil {
		return
	}

	resp = toFindVehicleByIdAndTimeResponse(domain)

	return
}
