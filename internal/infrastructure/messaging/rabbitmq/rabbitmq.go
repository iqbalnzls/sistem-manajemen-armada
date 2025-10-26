package rabbitmq

import (
	"context"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"

	"github.com/iqbalnzls/sistem-manajemen-armada/internal/common/config"
	"github.com/iqbalnzls/sistem-manajemen-armada/internal/common/logger"
)

type Publisher interface {
	PublishEvent(ctx context.Context, message []byte, key string) (err error)
}

type publisher struct {
	channel *amqp.Channel
	cfg     *config.RabbitMqConfig
}

func NewRabbitMQ(channel *amqp.Channel, cfg *config.RabbitMqConfig) Publisher {
	if channel == nil {
		panic("channel is nil")
	}
	if cfg == nil {
		panic("cfg is nil")
	}

	return &publisher{
		channel: channel,
		cfg:     cfg,
	}
}

func (p *publisher) PublishEvent(ctx context.Context, message []byte, key string) (err error) {
	log := logger.FromContext(ctx)

	log.Info("Publishing geofence event",
		zap.String("tag", "T2"),
		zap.Int64("start_process", time.Now().UnixMilli()),
		zap.Any("message", string(message)),
	)

	err = p.channel.PublishWithContext(
		ctx,
		p.cfg.Exchange.Name,
		key,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        message,
			Timestamp:   time.Now(),
		},
	)
	if err != nil {
		log.Error("Failed to publish geofence event", zap.Error(err))
		return
	}

	log.Info("Publishing geofence event",
		zap.String("tag", "T3"),
		zap.Int64("end_process", time.Now().UnixMilli()),
	)

	return
}
