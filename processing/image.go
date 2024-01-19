package processing

import (
	"fmt"
	"math"
	"os"
	"path/filepath"

	"github.com/h2non/bimg"
)

func ProcessImage(uploadedImage []byte, name string) error {
	conf := GetConfigurations()

	newImage := bimg.NewImage(uploadedImage)
	for _, c := range conf {
		img := set_image_scale(*newImage, c)
		save_with_format("./", c.name, name, img, bimg.JPEG)
		save_with_format("./", c.name, name, img, bimg.WEBP)
	}
	return nil
}

func set_image_scale(image bimg.Image, conf Image_config) *bimg.Image {

	size, err := image.Size()
	if err != nil {
		fmt.Println("Error determining image size")
	}
	scaler := max(size.Height, size.Width)
	fmt.Println(scaler)
	factor := float64(conf.height) / float64(scaler)
	fmt.Println(factor)
	resizedImg, err := image.Resize(int(math.Round(float64(size.Width)*factor)), int(math.Round(float64(size.Height)*factor)))

	return bimg.NewImage(resizedImg)

}

func save_with_format(path string, size string, name string, image *bimg.Image, format bimg.ImageType) error {
	options := bimg.Options{
		Quality: 100,
		Type:    format,
	}
	converted_img, err := image.Process(options)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	if err := os.WriteFile(path+name+"-"+size+"."+bimg.ImageTypes[format], converted_img, 0604); err != nil {
		return err
	}
	return nil
}
