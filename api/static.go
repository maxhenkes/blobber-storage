package api

import (
	"os"

	"github.com/gin-gonic/gin"
)

func EnableStaticRoute(r *gin.Engine) {
	path := os.Getenv("PATH_STORAGE")
	r.Static("/images", path)
}
