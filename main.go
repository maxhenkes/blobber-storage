package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/maxhenkes/blobber-storage/api"
)

func main() {

	api.StartWebserver()

}
