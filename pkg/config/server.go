package config

import (
	"context"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

func NewServer(lc fx.Lifecycle, c *Config) *fiber.App {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		DisableKeepalive:      true,
	})

	serverLifecycle(lc, c, app)

	return app
}

func serverLifecycle(lc fx.Lifecycle, c *Config, app *fiber.App) {

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				log.Info("Starting server...")
				err := app.Listen(":" + c.Port)
				if err != nil {
					log.Fatal(err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Info("Stopping server...")
			_ = app.Shutdown()
			log.Info("Server stopped")
			return nil
		},
	})
}
