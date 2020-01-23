package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"sync"

	"github.com/tazzaoui/risearch/lib"
	"gocv.io/x/gocv"
)

const MAX_GO_ROUTINES = 100
const MAX_IMAGES = -1 // Number of images to index (-1 for entire dir)

func main() {
	image_dir := "data/img"
	kp_dir := "data/kp"
	var wg sync.WaitGroup

	images, err := ioutil.ReadDir(image_dir)

	if err != nil {
		fmt.Println("Please download the image data set by running get_data.sh")
		os.Exit(1)
	}

	sem := make(chan int, MAX_GO_ROUTINES)

	i := 0
	for _, img := range images {
		if MAX_IMAGES > 0 && i >= MAX_IMAGES {
			break
		}

		img_path := path.Join(image_dir, img.Name())
		desc_path := path.Join(kp_dir, img.Name()+".tiff")

		sem <- 1
		go func() {
			wg.Add(1)
			desc := lib.GetDescriptors(img_path)

			// Dump descriptors to disk
			gocv.IMWrite(desc_path, desc)
			desc.Close()
			wg.Done()

			<-sem
		}()
		i++
	}

	wg.Wait()
}
