package delivery

import (
	"fmt"

	"github.com/iqbalnzls/sistem-manajemen-armada/internal/container"
	"github.com/iqbalnzls/sistem-manajemen-armada/internal/delivery/messaging"
	"github.com/iqbalnzls/sistem-manajemen-armada/internal/delivery/rest"
)

func StartServer(c *container.Container) {
	// Start HTTP server in background
	go rest.StartHttpServer(c)

	// Start MQTT messaging server (non-blocking)
	messaging.StartMessagingServer(c)

	fmt.Println("✅ All servers started successfully")
}
