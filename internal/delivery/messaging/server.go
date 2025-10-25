package messaging

import (
	"go.uber.org/zap"

	"github.com/iqbalnzls/sistem-manajemen-armada/internal/container"
)

func StartMessagingServer(c *container.Container) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	logger = logger.With(
		zap.String("service", c.Config.App.Name),
		zap.String("mqtt_topic", c.Config.MQTT.TopicPattern),
	)

	h := NewMessagingHandler(c, logger)

	if err := c.Mqtt.Subscribe(c.Config.MQTT.TopicPattern, h.vehicleLoc.ReceiveVehicleLocation); err != nil {
		panic(err)
	}
}
