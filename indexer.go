package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/tazzaoui/risearch/lib"
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
		kp_path := path.Join(kp_dir, img.Name()+".json")

		sem <- 1
		go func() {
			desc := lib.GetDescriptors(img_path)

			// Dump descriptors to disk
			bytes, _ := json.Marshal(desc)
			ioutil.WriteFile(kp_path, bytes, 0644)

		}()
		<-sem
	}
}
