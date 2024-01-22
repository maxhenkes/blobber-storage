package processing

import (
	"fmt"
	"log"
	"math"

	"github.com/h2non/bimg"
)

func ProcessImage(uploadedImage []byte, hash string) error {
	isProcessed := doesFileExist(hash, "thumb", "jpeg")

	if isProcessed {
		log.Printf("Image %s already processed, skipping...", hash)
		return nil
	}

	conf := GetConfigurations()

	newImage := bimg.NewImage(uploadedImage)
	for _, c := range conf {
		img := set_image_scale(*newImage, c)
		SaveWithFormat(c.name, hash, img, bimg.JPEG)
		SaveWithFormat(c.name, hash, img, bimg.WEBP)
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
