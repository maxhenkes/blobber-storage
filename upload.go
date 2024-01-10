package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var startTime = time.Now()

func enableUploadRoute() {
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/status", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"version": "0.1.0-SNAPSHOT",
			"uptime:": time.Since(startTime).Truncate(time.Second).String(),
		})
	})
	r.MaxMultipartMemory = 128 << 20
	r.POST("/upload", func(c *gin.Context) {
		//single File
		file, _ := c.FormFile("file")
		form, _ := c.MultipartForm()
		name := form.Value["name"]
		openedFile, _ := file.Open()
		rawFile, _ := io.ReadAll(openedFile)
		ProcessImage(rawFile)
		log.Println(file.Filename, file.Size, name)
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
	r.Run()
	//http.ListenAndServe(":8080", r) // listen and serve on 0.0.0.0:8080
}
