package rabbitmq

type EventType string

const (
	EventTypeGeofenceEntry EventType = "geofence.entry"
)

func (e EventType) String() string {
	return string(e)
}

type GeofenceEvent struct {
	VehicleID string    `json:"vehicle_id"`
	EventType EventType `json:"event"`
	Location  Location  `json:"location"`
	Timestamp int64     `json:"timestamp"`
}

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
