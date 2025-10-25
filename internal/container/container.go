package container

import (
	"github.com/iqbalnzls/sistem-manajemen-armada/internal/app/usecase/vehiclelocations"
	"github.com/iqbalnzls/sistem-manajemen-armada/internal/infrastructure/messaging/rabbitmq"
	vehicleLocationsRepo "github.com/iqbalnzls/sistem-manajemen-armada/internal/infrastructure/postgres/vehicle_locations"
	"github.com/iqbalnzls/sistem-manajemen-armada/pkg/config"
	"github.com/iqbalnzls/sistem-manajemen-armada/pkg/database"
	"github.com/iqbalnzls/sistem-manajemen-armada/pkg/messaging"
	"github.com/iqbalnzls/sistem-manajemen-armada/pkg/validator"
)

type Container struct {
	Config        *config.Config
	VehicleLocSvc vehiclelocations.Service
	Validator     *validator.Validator
	RabbitMQ      *messaging.RabbitMQ
	Mqtt          *messaging.MQTT
}

func SetupContainer() *Container {
	//init config
	cfg := config.NewConfig("resources/config.json")

	//init validator
	v := validator.NewValidator()

	//init rabbitmq
	rabbitMq := messaging.NewRabbitMQConnection(&cfg.RabbitMQ)

	//init mqtt
	mqtt := messaging.NewMQTTConnection(&cfg.MQTT)

	//init database
	db := database.NewDatabase(&cfg.Database)

	//init infrastructure
	vehicleLocRepo := vehicleLocationsRepo.NewVehicleLocationsRepository(db)
	rabbitmqPublisher := rabbitmq.NewRabbitMQ(rabbitMq.Channel, &cfg.RabbitMQ)

	//init service
	vehicleLocSvc := vehiclelocations.NewVehicleLocationsService(vehicleLocRepo, rabbitmqPublisher)

	return &Container{
		Config:        cfg,
		VehicleLocSvc: vehicleLocSvc,
		Validator:     v,
		RabbitMQ:      rabbitMq,
		Mqtt:          mqtt,
	}
}

func (c *Container) Cleanup() {
	// Close MQTT connection
	c.Mqtt.Close()

	// Close RabbitMQ connections
	c.RabbitMQ.Channel.Close()
	c.RabbitMQ.Conn.Close()
}
