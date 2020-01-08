package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"gocv.io/x/gocv"
	"gocv.io/x/gocv/contrib"
)

func main() {
	image_dir := "data/img"

	images, err := ioutil.ReadDir(image_dir)

	if err != nil {
		fmt.Println("Please download the image data set by running get_data.sh")
		os.Exit(1)
	}

	for _, img := range images {
		fmt.Println(img.Name())

		img := gocv.IMRead(img.Name(), gocv.IMReadGrayScale)
		defer img.Close()

		sift := contrib.NewSIFT()
		defer sift.Close()

		key_points := sift.Detect(img)
		fmt.Println(key_points)
	}
}
