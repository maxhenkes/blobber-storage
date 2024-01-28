package api

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

func StartWebserver() {
	startFiber()
}

func startGin() {
	r := gin.Default()
	r.Use(CORS())
	EnableUploadRoute(r)
	EnableStaticRoute(r)
	EnableHealthRoute(r)
	//r.Run("127.0.0.1:4545")
	r.Run()
	//http.ListenAndServe(":8080", r) // listen and serve on 0.0.0.0:8080
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
	EnableUploadRouteF(app)
	app.Get("/", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{"test": "test"})
	})

	app.Listen(":3010")

}

func CORS() gin.HandlerFunc {
	// TO allow CORS
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, Token")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func TokenAuth() gin.HandlerFunc {
	token := os.Getenv("ACCESS_TOKEN")
	log.Println("Hitting middleware")

	if token == "" {
		log.Fatal("Please set ACCESS_TOKEN environment variable")
	}
	return func(c *gin.Context) {

		requestToken := c.Request.Header.Get("Token")

		if requestToken == "" {
			respondWithError(c, 401, "API token required")
			return
		}

		if requestToken != token {
			log.Println("Token doesn't match")
			respondWithError(c, 401, "Invalid API token")
			return
		}
		log.Println("test")

		c.Next()
	}
}

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}
