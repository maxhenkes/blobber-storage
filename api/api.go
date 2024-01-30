package api

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

func StartWebserver() {
	startFiber()
}

func startFiber() {
	app := fiber.New()
	app.Use(logger.New())
	app.Server().MaxRequestBodySize = 100 * 1024 * 1024
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "Origin, Content-Type, Accept, Token",
	}))
	app.Use("/upload", func(c fiber.Ctx) error {
		fmt.Println("Token info: ")
		fmt.Println(c.Get("Token"))
		reqToken := c.Get("Token")

		if reqToken == "" {
			return c.SendStatus(401)
		}
		if reqToken != "test-1234-token" {
			return c.SendStatus(401)
		}

		return c.Next()
	})
	EnableUploadRoute(app)
	EnableStaticRoute(app)
	EnableHealthRoute(app)
	app.Get("/", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{"test": "test"})
	})

	app.Listen(":3010")

}
