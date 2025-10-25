package dto

type ReceiveVehicleLocation struct {
	VehicleId string  `json:"vehicle_id" validate:"required,platenumber"`
	Latitude  float64 `json:"latitude" validate:"required,latitude"`
	Longitude float64 `json:"longitude" validate:"required,longitude"`
	Timestamp int64   `json:"timestamp" validate:"gt=0"`
}

type FindVehicleByIdRequest struct {
	VehicleId string `validate:"required,platenumber"`
}

type FindVehicleByIdResponse struct {
	FindVehicleResponse
}

type FindVehicleByIdAndTimeRequest struct {
	VehicleId string `validate:"required,platenumber"`
	Start     int64  `query:"start" validate:"gt=0"`
	End       int64  `query:"end" validate:"gt=0"`
}

type FindVehicleByIdAndTimeResponse struct {
	FindVehicleResponse
}

type FindVehicleResponse struct {
	VehicleId string  `json:"vehicle_id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Timestamp int64   `json:"timestamp"`
}
