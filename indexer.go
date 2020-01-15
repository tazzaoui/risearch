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

	sem := make(chan int, MAX_GO_ROUTINES)

	for _, img := range images {
		img_path := path.Join(image_dir, img.Name())
		kp_path := path.Join(kp_dir, img.Name()+".json")

		sem <- 1
		go func() {
			mat := gocv.IMRead(img_path, gocv.IMReadGrayScale)
			defer mat.Close()

			// Extract descriptors
			if !mat.Empty() {
				mask := gocv.NewMat()
				key_pts, desc := sift.DetectAndCompute(mat, mask)

				var descriptors [][]float64
				for i, _ := range key_pts {
					var tmp []float64
					for j := 0; j < desc.Cols(); j++ {
						tmp = append(tmp, float64(desc.GetFloatAt(i, j)))
					}
					descriptors = append(descriptors, tmp)
				}

				// Dump descriptors to disk
				bytes, _ := json.Marshal(descriptors)
				ioutil.WriteFile(kp_path, bytes, 0644)

				mat.Close()
				desc.Close()
			}
			<-sem
		}()
	}
}
