package main

import (
	"fmt"
	"github.com/tazzaoui/risearch/lib"
	"gocv.io/x/gocv"
	"io/ioutil"
	"os"
	"path"
)

func read_descs(max int) (descs []gocv.Mat) {
	kp_dir := "data/kp"
	files, err := ioutil.ReadDir(kp_dir)

	if err != nil {
		fmt.Println("Keypoints should be in data/kp")
		os.Exit(1)
	}

	i := 0
	for _, f := range files {
		if max > 0 && i >= max {
			break
		}
		kp_path := path.Join(kp_dir, f.Name())
		desc := gocv.IMRead(kp_path, gocv.IMReadUnchanged)

		descs = append(descs, desc)
		i++
	}
	return descs
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Please provide an input image.")
		return
	}

	img_desc := lib.GetDescriptors(os.Args[1])

	bf := gocv.NewBFMatcherWithParams(gocv.NormL1, false)
	defer bf.Close()

	desc_db := read_descs(100)

	for _, d := range desc_db {
		matches := bf.KnnMatch(d, img_desc, 4)
		fmt.Println(matches)
	}
}
