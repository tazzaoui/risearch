package main

import (
	"fmt"
	"os"

	"gocv.io/x/gocv"
	"gocv.io/x/gocv/contrib"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Please provide an input image.")
		return
	}

	img_path := os.Args[1]

	sift := contrib.NewSIFT()
	defer sift.Close()

	mat := gocv.IMRead(img_path, gocv.IMReadGrayScale)
	defer mat.Close()

	if !mat.Empty() {
		key_pts := sift.Detect(mat)
		fmt.Println(key_pts)
	}
}
