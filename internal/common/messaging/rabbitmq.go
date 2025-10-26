package messaging

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/iqbalnzls/sistem-manajemen-armada/internal/common/config"
)

type RabbitMQ struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

func NewRabbitMQConnection(cfg *config.RabbitMqConfig) *RabbitMQ {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%v/", cfg.User, cfg.Pass, cfg.Host, cfg.Port))
	if err != nil {
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		panic(err)
	}

	err = ch.ExchangeDeclare(
		cfg.Exchange.Name,
		cfg.Exchange.Type,
		cfg.Exchange.Durable,
		cfg.Exchange.AutoDelete,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}

	queue, err := ch.QueueDeclare(
		cfg.Queue.Name,
		cfg.Queue.Durable,
		cfg.Queue.AutoDelete,
		cfg.Queue.Exclusive,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}

	err = ch.QueueBind(
		queue.Name,
		cfg.Queue.RoutingKey,
		cfg.Exchange.Name,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}

	fmt.Println("✅ RabbitMQ initialized successfully")

	return &RabbitMQ{
		Conn:    conn,
		Channel: ch,
	}
}

//func (c *RabbitMQ) ConsumeGeofenceAlerts(queueName string) error {
//	msgs, err := c.Channel.Consume(
//		queueName,
//		"",
//		false,
//		false,
//		false,
//		false,
//		nil,
//	)
//	if err != nil {
//		log.Println("Failed to start consuming messages:")
//		return err
//	}
//
//	go func() {
//		for msg := range msgs {
//			log.Println("✅ Received RabbitMQ message:", string(msg.Body))
//		}
//	}()
//
//	return nil
//}
