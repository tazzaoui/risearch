package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

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

	sift := contrib.NewSIFT()
	defer sift.Close()

	for _, img := range images {
		img_path := path.Join(image_dir, img.Name())

		mat := gocv.IMRead(img_path, gocv.IMReadGrayScale)
		defer mat.Close()

		if !mat.Empty() {
			key_points := sift.Detect(mat)
			fmt.Println(key_points)
		}
	}
}
