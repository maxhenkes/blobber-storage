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

func doesFileExist(hash string, size string, format string) bool {
	path := os.Getenv("PATH_STORAGE")
	filePath := fmt.Sprintf("%s%s", path, hash)
	_, err := os.Stat(filePath)
	return os.IsNotExist(err)
}
