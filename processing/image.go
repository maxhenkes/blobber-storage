package processing

import (
	"fmt"
	"log"
	"math"

	"github.com/h2non/bimg"
)

type Image struct {
	Data []byte
	Hash string
}

func ProcessImage(image Image) error {
	isProcessed := doesFileExist(image.Hash, "thumb", "jpeg")

	if isProcessed {
		log.Printf("Image %s already processed, skipping...", image.Hash)
		return nil
	}

	conf := GetConfigurations()

	newImage := bimg.NewImage(image.Data)
	for _, c := range conf {
		img := set_image_scale(*newImage, c)
		SaveWithFormat(c.name, image.Hash, img, bimg.JPEG)
		SaveWithFormat(c.name, image.Hash, img, bimg.WEBP)
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
