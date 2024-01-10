package main

import (
	"fmt"
	"os"

	"github.com/h2non/bimg"
)

func ProcessImage(uploadedImage []byte) {

	newImage, err := bimg.NewImage(uploadedImage).Convert(bimg.WEBP)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	bimg.Write("new.webp", newImage)
}
