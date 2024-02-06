package processing

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/h2non/bimg"
)

func SaveWithFormat(size string, name string, image *bimg.Image, format bimg.ImageType) error {
	path := os.Getenv("PATH_STORAGE")
	options := bimg.Options{
		Quality: 100,
		Type:    format,
	}
	converted_img, err := image.Process(options)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(path+"/"+name+"/"), 0755); err != nil {
		return err
	}
	if err := os.WriteFile(path+"/"+name+"/"+name+"-"+size+"."+bimg.ImageTypes[format], converted_img, 0604); err != nil {
		return err
	}
	return nil
}

func doesFileExist(hash string) bool {
	path := os.Getenv("PATH_STORAGE")
	filePath := fmt.Sprintf("%s%s", path, hash)
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}

func CheckAndReturnConfig(hash string) Config {
	path := os.Getenv("PATH_STORAGE")
	filePath := fmt.Sprintf("%s%s", path, hash)
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return Config{}
	}

	files, err := os.ReadDir(fmt.Sprintf("%s%s/", path, hash))
	if err != nil {
		return Config{}
	}
	bConfig := config.Configs
	aConfig := []Image_config{}

	for _, conf := range bConfig {
		found := false
		for _, file := range files {

			if file.Name() == fmt.Sprintf("%s-%s.jpeg", hash, conf.Name) {
				found = true
				break
			}

		}
		if !found {
			aConfig = append(aConfig, conf)
		}
	}
	return Config{aConfig}
}
