package api

import (
	"fmt"
	"os"

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

	EnableUploadRoute(app)
	EnableStaticRoute(app)
	EnableHealthRoute(app)
	EnableManageRoutes(app)

	app.Listen(":3010")

}

func TokenMiddleware(c fiber.Ctx) error {
	fmt.Println("Token info: ")
	fmt.Println(c.Get("Token"))
	appToken := os.Getenv("ACCESS_TOKEN")
	reqToken := c.Get("Token")

	if reqToken == "" {
		return c.SendStatus(401)
	}
	if reqToken != appToken {
		return c.SendStatus(401)
	}

	return c.Next()
}
