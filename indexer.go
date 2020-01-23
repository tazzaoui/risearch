package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/tazzaoui/risearch/lib"
	"gocv.io/x/gocv"
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

	sem := make(chan int, MAX_GO_ROUTINES)

	for _, img := range images {
		img_path := path.Join(image_dir, img.Name())
		desc_path := path.Join(kp_dir, img.Name()+".png")

		sem <- 1
		go func() {
			desc := lib.GetDescriptors(img_path)

			// Dump descriptors to disk
			gocv.IMWrite(desc_path, desc)
			desc.Close()

			<-sem
		}()

	}
}
