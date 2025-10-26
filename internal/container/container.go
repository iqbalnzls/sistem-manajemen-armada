package container

import (
	"os"

	"github.com/iqbalnzls/sistem-manajemen-armada/internal/common/config"
	"github.com/iqbalnzls/sistem-manajemen-armada/internal/common/database"
	messaging2 "github.com/iqbalnzls/sistem-manajemen-armada/internal/common/messaging"
	"github.com/iqbalnzls/sistem-manajemen-armada/internal/common/validator"
	"github.com/iqbalnzls/sistem-manajemen-armada/internal/infrastructure/messaging/rabbitmq"
	vehicleLocationsRepo "github.com/iqbalnzls/sistem-manajemen-armada/internal/infrastructure/postgres/vehicle_locations"
	vehiclelocations2 "github.com/iqbalnzls/sistem-manajemen-armada/internal/usecase/vehiclelocations"
)

type Container struct {
	Config        *config.Config
	VehicleLocSvc vehiclelocations2.Service
	Validator     *validator.Validator
	RabbitMQ      *messaging2.RabbitMQ
	Mqtt          *messaging2.MQTT
}

func SetupContainer() *Container {
	//init config
	cfg := config.NewConfig(os.Getenv("CONFIG_FILE"))

	//init validator
	v := validator.NewValidator()

	//init rabbitmq
	rabbitMq := messaging2.NewRabbitMQConnection(&cfg.RabbitMQ)

	//init mqtt
	mqtt := messaging2.NewMQTTConnection(&cfg.MQTT)

	//init database
	db := database.NewDatabase(&cfg.Database)

	//init infrastructure
	vehicleLocRepo := vehicleLocationsRepo.NewVehicleLocationsRepository(db)
	rabbitmqPublisher := rabbitmq.NewRabbitMQ(rabbitMq.Channel, &cfg.RabbitMQ)

	//init service
	vehicleLocSvc := vehiclelocations2.NewVehicleLocationsService(vehicleLocRepo, rabbitmqPublisher)

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
