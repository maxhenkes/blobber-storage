package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StartWebserver() {
	r := gin.Default()
	r.Use(cors.Default())
	EnableUploadRoute(r)
	EnableStaticRoute(r)
	EnableHealthRoute(r)
	//r.Run("127.0.0.1:4545")
	r.Run()
	//http.ListenAndServe(":8080", r) // listen and serve on 0.0.0.0:8080
}
