package processing

import (
	"errors"
	"fmt"
	"math"

	"github.com/h2non/bimg"
)

type Image struct {
	Data []byte
	Hash string
}

func ProcessImage(image Image) error {
	if len(config.Configs) == 0 {
		return errors.New("cannot check if file exists because config is empty")
	}

	leftConfig := CheckAndReturnConfig(image.Hash)

	newImage := bimg.NewImage(image.Data)
	for _, c := range leftConfig.Configs {
		img := set_image_scale(*newImage, c)
		SaveWithFormat(c.Name, image.Hash, img, bimg.JPEG)
		SaveWithFormat(c.Name, image.Hash, img, bimg.WEBP)
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
	factor := float64(conf.Height) / float64(scaler)
	fmt.Println(factor)
	resizedImg, _ := image.Resize(int(math.Round(float64(size.Width)*factor)), int(math.Round(float64(size.Height)*factor)))

	return bimg.NewImage(resizedImg)

}
