package api

import (
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/maxhenkes/blobber-storage/util"
)

var formats = []string{"large.jpeg", "large.webp", "medium.jpeg", "medium.webp", "small.jpeg", "small.webp", "thumb.jpeg", "thumb.webp"}

func EnableManageRoutes(app *fiber.App) {
	app.Get("/manage/all", func(c fiber.Ctx) error {
		includeFiles := c.Query("includeFiles")
		path := os.Getenv("PATH_STORAGE")
		files, err := os.ReadDir(path)
		if err != nil {
			return c.SendStatus(500)
		}

		if includeFiles != "true" {
			return c.JSON(util.Map(files, func(entry os.DirEntry) string {
				return entry.Name()
			}))
		}

		type jsonResponse struct {
			Id     string
			Images []string
		}

		fileList := util.Map(files, func(entry os.DirEntry) jsonResponse {
			imageList := []string{}

			for _, v := range formats {
				imageList = append(imageList, entry.Name()+"-"+v)
			}

			return jsonResponse{Id: entry.Name(), Images: imageList}
		})

		return c.JSON(fileList)
	}, TokenMiddleware)

	app.Delete("/manage/images/:id", func(c fiber.Ctx) error {
		path := os.Getenv("PATH_STORAGE")
		id := c.Params("id")
		log.Println(id)
		if strings.TrimSpace(id) == "" {
			return c.SendStatus(422)
		}

		pathIdString := path + id

		_, err := os.ReadDir(pathIdString)
		if err != nil {
			return c.SendStatus(404)
		}

		pathError := os.RemoveAll(pathIdString)
		if pathError != nil {
			return c.SendStatus(500)
		}
		return c.JSON(id)
	}, TokenMiddleware)
}
