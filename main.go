package main

import (
	"fmt"
	"github.com/tazzaoui/risearch/lib"
	"gocv.io/x/gocv"
	"io/ioutil"
	"os"
	"path"
)

type match struct {
	img     string          // Path to image being matched against
	matches [][]gocv.DMatch // Matches ranked in increasing order by distance
}

const MAX_IMAGES = 100

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Please provide an input image.")
		return
	}

	kp_dir := "data/kp"
	files, err := ioutil.ReadDir(kp_dir)

	if err != nil {
		fmt.Println("Keypoints should be in data/kp")
		os.Exit(1)
	}

	img_desc := lib.GetDescriptors(os.Args[1])

	bf := gocv.NewBFMatcherWithParams(gocv.NormL1, false)
	defer bf.Close()

	var matches []match

	i := 0
	for _, f := range files {
		if MAX_IMAGES > 0 && i >= MAX_IMAGES {
			break
		}

		kp_path := path.Join(kp_dir, f.Name())
		desc := gocv.IMRead(kp_path, gocv.IMReadUnchanged)

		desc_matches := bf.KnnMatch(desc, img_desc, 4)
		matches = append(matches, match{kp_path, desc_matches})
		i++
	}

	fmt.Println(matches)
}
