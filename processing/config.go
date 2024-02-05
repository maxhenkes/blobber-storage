package processing

import (
	"encoding/json"
	"fmt"
	"os"
)

var config Config

type Config struct {
	Configs []Image_config `json:"configs"`
}

type Image_config struct {
	Height int    `json:"height"`
	Width  int    `json:"width"`
	Name   string `json:"name"`
}

func LoadConfiguration(file string) {
	fmt.Println("Loading configuration file...")
	var localConfig Config
	configFile, err := os.Open(file)
	if err != nil {
		fmt.Println("Error loading configuration file...")
		fmt.Println(err.Error())
		configFile.Close()
		panic("Cannot start without configuration file...")
	}
	defer configFile.Close()
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&localConfig)
	config = localConfig
}
