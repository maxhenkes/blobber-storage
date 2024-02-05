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
		file, err := c.FormFile("file")
		if err != nil {
			return c.SendStatus(500)
		}
		form, err := c.MultipartForm()
		if err != nil {
			return c.SendStatus(500)
		}
		formName := form.Value["name"]
		if len(formName) != 1 {
			return c.SendStatus(500)
		}
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
		log.Println(file.Filename, file.Size, formName[0])
		return c.Status(fiber.StatusAccepted).JSON(&fiber.Map{"data": hash})
	}, TokenMiddleware)
}
