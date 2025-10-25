package messaging

import (
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"

	"github.com/iqbalnzls/sistem-manajemen-armada/pkg/config"
)

type MQTT struct {
	Client mqtt.Client
}

type MessageHandler func(topic string, payload []byte) error

func NewMQTTConnection(cfg *config.MQTTConfig) *MQTT {
	opts := mqtt.NewClientOptions()
	brokerURL := fmt.Sprintf("tcp://%s:%d", cfg.Broker, cfg.Port)
	opts.AddBroker(brokerURL)
	opts.SetClientID(cfg.ClientID)
	opts.SetUsername(cfg.Username)
	opts.SetPassword(cfg.Password)
	opts.SetCleanSession(cfg.CleanSession)
	opts.SetAutoReconnect(cfg.AutoReconnect)
	opts.SetConnectRetry(true)
	opts.SetConnectRetryInterval(time.Duration(cfg.ConnectRetryDelay) * time.Second)
	opts.SetMaxReconnectInterval(1 * time.Minute)

	opts.OnConnect = func(c mqtt.Client) {
		fmt.Println(fmt.Sprintf("[MQTT] Connected to broker:%s", brokerURL))
	}

	opts.OnConnectionLost = func(c mqtt.Client, err error) {
		fmt.Println(fmt.Sprintf("[MQTT] Connection lost:%v", err))
	}

	client := mqtt.NewClient(opts)

	token := client.Connect()
	token.Wait()

	if err := token.Error(); err != nil {
		log.Printf("[MQTT] Failed to connect: %v", err)
		panic(err)
	}

	// Small delay to ensure connection is fully established
	// The OnConnect callback fires slightly before IsConnected() returns true
	time.Sleep(100 * time.Millisecond)

	fmt.Println("✅ MQTT connection established successfully")

	return &MQTT{
		Client: client,
	}
}

func (m *MQTT) Subscribe(topic string, handler MessageHandler) error {
	messageHandler := func(client mqtt.Client, msg mqtt.Message) {
		if err := handler(msg.Topic(), msg.Payload()); err != nil {
			log.Printf("[MQTT] Error handling message on topic %s: %v", msg.Topic(), err)
		}
	}

	token := m.Client.Subscribe(topic, 1, messageHandler)
	token.Wait()

	if err := token.Error(); err != nil {
		return err
	}

	return nil
}

func (m *MQTT) Close() {
	log.Println("[MQTT] Closing connection...")

	// Wait 250ms for pending messages
	m.Client.Disconnect(250)

	log.Println("[MQTT] ✅ Connection closed")
}
