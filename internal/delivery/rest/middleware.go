package rest

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/iqbalnzls/sistem-manajemen-armada/internal/common/config"
	"github.com/iqbalnzls/sistem-manajemen-armada/internal/common/logger"
)

func SetupMiddleware(log *zap.Logger, cfg *config.AppConfig) fiber.Handler {
	return func(c *fiber.Ctx) error {
		log = log.With(
			zap.String("service", cfg.Name),
			zap.String("xid", uuid.New().String()),
			zap.String("uri", c.Request().URI().String()),
			zap.String("method", c.Method()),
		)

		tn := time.Now()
		log.Info("Incoming request",
			zap.Any("req", c.Body()),
			zap.String("tag", "T1"),
			zap.String("timestamp", tn.Format(time.RFC3339)),
		)

		ctx := c.UserContext()
		ctx = logger.WithLogger(ctx, log)
		ctx = context.WithValue(ctx, "startTime", tn)
		c.SetUserContext(ctx)

		return c.Next()
	}
}
