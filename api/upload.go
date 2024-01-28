package api

import (
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v3"
	"github.com/maxhenkes/blobber-storage/processing"
	"github.com/maxhenkes/blobber-storage/util"
)

func EnableUploadRoute(r *gin.Engine) {
	r.MaxMultipartMemory = 128 << 20

	r.POST("/upload", TokenAuth(), func(c *gin.Context) {
		log.Printf("UPLOAD ROUTE")
		//single File
		file, _ := c.FormFile("file")
		form, _ := c.MultipartForm()
		name := form.Value["name"]
		openedFile, _ := file.Open()
		rawFile, _ := io.ReadAll(openedFile)

		hash := util.ComputeHashFromFile(&rawFile)

		go processing.ProcessImage(rawFile, hash)

		log.Println(file.Filename, file.Size, name)
		c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "", "data": hash})
	})

}

func EnableUploadRouteF(app *fiber.App) {
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
		go processing.ProcessImage(rawFile, hash)
		log.Println(file.Filename, file.Size, name)
		return c.Status(fiber.StatusAccepted).JSON(&fiber.Map{"data": hash})
	})
}
