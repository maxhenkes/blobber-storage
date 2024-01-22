package api

import (
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maxhenkes/blobber-storage/processing"
	"github.com/maxhenkes/blobber-storage/util"
)

func EnableUploadRoute(r *gin.Engine) {
	r.MaxMultipartMemory = 128 << 20
	r.POST("/upload", func(c *gin.Context) {
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
