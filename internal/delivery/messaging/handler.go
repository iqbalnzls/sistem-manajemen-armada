package messaging

type VehicleLocationHandler interface {
	ReceiveVehicleLocation(topic string, payload []byte) error
}
