package api

import (
	"io"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/maxhenkes/blobber-storage/processing"
	"github.com/maxhenkes/blobber-storage/util"
)

func EnableUploadRoute(app *fiber.App) {
	app.Post("/upload", func(c fiber.Ctx) error {
		log.Println("Route")
		file, err := c.FormFile("file")
		if err != nil {
			return c.SendStatus(500)
		}
		form, err := c.MultipartForm()
		if err != nil {
			return c.SendStatus(500)
		}
		name := form.Value["name"][0]
		openedFile, err := file.Open()
		if err != nil {
			return c.SendStatus(500)
		}
		rawFile, err := io.ReadAll(openedFile)
		if err != nil {
			return c.SendStatus(500)
		}

		hash := util.ComputeHashFromFile(&rawFile)

		image := processing.Image{Data: rawFile, Hash: hash}

		go processing.ProcessImage(image)
		log.Println(file.Filename, file.Size, name)
		return c.Status(fiber.StatusAccepted).JSON(&fiber.Map{"data": hash})
	}, TokenMiddleware)
}
