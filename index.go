package main

import (
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

		c <- 1
		go func() {
			mat := gocv.IMRead(img_path, gocv.IMReadGrayScale)
			defer mat.Close()

			if !mat.Empty() {
				key_pts, desc := sift.DetectAndCompute(mat, gocv.NewMat())
				var descriptors [][]float64
				for i, _ := range key_pts {
					var tmp []float64
					for j := 0; j < desc.Cols(); j++ {
						tmp = append(tmp, float64(desc.GetFloatAt(i, j)))
					}
					descriptors = append(descriptors, tmp)
				}
				fmt.Println(descriptors)
			}
			<-c
		}()
	}
}
