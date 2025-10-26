package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.uber.org/zap"

	"github.com/iqbalnzls/sistem-manajemen-armada/internal/common/logger"
	"github.com/iqbalnzls/sistem-manajemen-armada/internal/container"
)

func StartHttpServer(container *container.Container) {
	// Initialize HTTP-specific logger
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	app := fiber.New(fiber.Config{
		ErrorHandler: errorHandler(),
		AppName:      container.Config.App.Name,
	})

	app.Use(SetupMiddleware(logger, &container.Config.App))
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	SetupRouter(app, NewRestHandler(container))

	_ = app.Listen(container.Config.AppAddress())
}

func errorHandler() fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		log := logger.FromContext(ctx.UserContext())

		log.Info("Outgoing request")

		return ctx.Status(fiber.StatusOK).JSON(
			fiber.Map{
				"result":  false,
				"message": err.Error(),
			})
	}
}
