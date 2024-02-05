package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/maxhenkes/blobber-storage/api"
	"github.com/maxhenkes/blobber-storage/processing"
)

func main() {
	processing.LoadConfiguration("config/config.json")
	api.StartWebserver()
}
