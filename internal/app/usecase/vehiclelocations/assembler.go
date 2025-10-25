package vehiclelocations

import (
	"github.com/iqbalnzls/sistem-manajemen-armada/internal/domain"
	"github.com/iqbalnzls/sistem-manajemen-armada/internal/dto"
	"github.com/iqbalnzls/sistem-manajemen-armada/internal/infrastructure/messaging/rabbitmq"
)

func toVehicleLocations(d *dto.ReceiveVehicleLocation) *domain.VehicleLocations {
	return &domain.VehicleLocations{
		VehicleId: d.VehicleId,
		Latitude:  d.Latitude,
		Longitude: d.Longitude,
		Timestamp: d.Timestamp,
	}
}

func toGeofenceEvent(d *dto.ReceiveVehicleLocation) *rabbitmq.GeofenceEvent {
	return &rabbitmq.GeofenceEvent{
		VehicleID: d.VehicleId,
		EventType: rabbitmq.EventTypeGeofenceEntry,
		Location: rabbitmq.Location{
			Latitude:  d.Latitude,
			Longitude: d.Longitude,
		},
		Timestamp: d.Timestamp,
	}
}

func toFindVehicleByIdResponse(domain *domain.VehicleLocations) dto.FindVehicleByIdResponse {
	return dto.FindVehicleByIdResponse{
		FindVehicleResponse: dto.FindVehicleResponse{
			VehicleId: domain.VehicleId,
			Latitude:  domain.Latitude,
			Longitude: domain.Longitude,
			Timestamp: domain.Timestamp,
		},
	}
}

func toFindVehicleByIdAndTimeResponse(domain *domain.VehicleLocations) dto.FindVehicleByIdAndTimeResponse {
	return dto.FindVehicleByIdAndTimeResponse{
		FindVehicleResponse: dto.FindVehicleResponse{
			VehicleId: domain.VehicleId,
			Latitude:  domain.Latitude,
			Longitude: domain.Longitude,
			Timestamp: domain.Timestamp,
		},
	}
}
