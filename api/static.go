package api

import (
	"os"

	"github.com/gofiber/fiber/v3"
)

func EnableStaticRoute(app *fiber.App) {
	path := os.Getenv("PATH_STORAGE")
	app.Static("/images", path)
}
