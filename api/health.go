package api

import (
	"time"

	"github.com/gofiber/fiber/v3"
)

var startTime = time.Now()

func EnableHealthRoute(app *fiber.App) {
	app.Get("/status", func(c fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"version": "0.1.0-SNAPSHOT",
			"uptime:": time.Since(startTime).Truncate(time.Second).String(),
		})
	})
}
