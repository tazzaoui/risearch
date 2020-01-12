package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"gocv.io/x/gocv"
	"gocv.io/x/gocv/contrib"
)

const MAX_GO_ROUTINES = 100

func main() {
	image_dir := "data/img"
	kp_dir := "data/kp"

	images, err := ioutil.ReadDir(image_dir)

	if err != nil {
		fmt.Println("Please download the image data set by running get_data.sh")
		os.Exit(1)
	}

	sift := contrib.NewSIFT()
	defer sift.Close()

	c := make(chan int, MAX_GO_ROUTINES)

	for _, img := range images {
		img_path := path.Join(image_dir, img.Name())
		kp_path := path.Join(kp_dir, img.Name()+".json")

		c <- 1
		go func() {
			mat := gocv.IMRead(img_path, gocv.IMReadGrayScale)
			defer mat.Close()

			if !mat.Empty() {
				key_pts, _ := sift.DetectAndCompute(mat, gocv.NewMat())
				encoded_kp, err := json.Marshal(key_pts)
				if err == nil {
					ioutil.WriteFile(kp_path, encoded_kp, 0644)
				}
			}
			<-c
		}()
	}
}
