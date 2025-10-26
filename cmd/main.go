package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/iqbalnzls/sistem-manajemen-armada/internal/container"
	"github.com/iqbalnzls/sistem-manajemen-armada/internal/delivery"
	"github.com/iqbalnzls/sistem-manajemen-armada/worker"
)

func main() {
	// Setup container
	c := container.SetupContainer()
	defer c.Cleanup()

	// Start all servers (HTTP + MQTT _+ RabbitMQ)
	delivery.StartServer(c)

	// Start RabbitMQ worker that consumes geofence_alerts queue
	worker.RabbitMQWorker(c.RabbitMQ, c.Config.RabbitMQ.Queue.Name)

	// Start MQTT scheduler to publish the vehicle location every 2 seconds
	worker.SchedulerMQTT(c.Mqtt)

	// Wait for the interrupt signal (Ctrl+C)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("Received shutdown signal, cleaning up...")
}
