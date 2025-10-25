package domain

type VehicleLocations struct {
	VehicleId string
	Latitude  float64
	Longitude float64
	Timestamp int64
}

func (VehicleLocations) TableName() string {
	return "vehicle_locations"
}
