package api

import (
	"time"

	"github.com/gin-gonic/gin"
)

var startTime = time.Now()

func EnableHealthRoute(r *gin.Engine) {
	r.GET("/status", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"version": "0.1.0-SNAPSHOT",
			"uptime:": time.Since(startTime).Truncate(time.Second).String(),
		})
	})

}
