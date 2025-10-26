package worker

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/iqbalnzls/sistem-manajemen-armada/internal/common/messaging"
)

const (
	JakartaLatitude  = -6.2088
	JakartaLongitude = 106.8325

	MaxLatOffset  = 0.00045
	MaxLongOffset = 0.00045
)

// RabbitMQWorker represent worker handles message consumption from the geofence_alerts queue.
func RabbitMQWorker(rmq *messaging.RabbitMQ, queueName string) {
	msgs, err := rmq.Channel.Consume(
		queueName,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Println("RabbitMQ failed to start consuming messages:")
		return
	}

	go func() {
		for msg := range msgs {
			log.Println("✅ RabbitMQ worker received message:", string(msg.Body))
		}
	}()
}

type VehicleLocationPayload struct {
	VehicleID string  `json:"vehicle_id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Timestamp int64   `json:"timestamp"`
}

// SchedulerMQTT publishes vehicle location events to MQTT every 2 seconds
// with randomly generated coordinates within ±50 meters of Jakarta
func SchedulerMQTT(mqtt *messaging.MQTT) {
	vehicleId := "B1234VV"
	topic := fmt.Sprintf("/fleet/vehicle/%s/location", vehicleId)

	log.Printf("🚀 Starting MQTT scheduler for vehicle %s on topic %s", vehicleId, topic)

	go func() {
		ticker := time.NewTicker(2 * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			// Generate random coordinates within ±50 meters of Jakarta
			lat := JakartaLatitude + (rand.Float64()*2-1)*MaxLatOffset
			long := JakartaLongitude + (rand.Float64()*2-1)*MaxLongOffset

			payload := VehicleLocationPayload{
				VehicleID: vehicleId,
				Latitude:  lat,
				Longitude: long,
				Timestamp: time.Now().Unix(),
			}

			message, _ := json.Marshal(payload)

			token := mqtt.Client.Publish(topic, 1, false, message)
			token.Wait()

			if token.Error() != nil {
				log.Printf("❌ MQTT Scheduler - Failed to publish to MQTT: %v", token.Error())
			} else {
				log.Printf("✅ MQTT Scheduler - Published to %s: %s", topic, string(message))
			}
		}
	}()
}
