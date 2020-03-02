package main

import (
	"fmt"
	"github.com/tazzaoui/risearch/lib"
	"gocv.io/x/gocv"
	"io/ioutil"
	"os"
	"path"
	"sort"
)

type match struct {
	img      string  // Path to image being matched against
	avg_dist float64 // Average Euclidean Distance between matches
}

const MAX_IMAGES = 100
const KP_DIR = "data/kp"
const K = 4

func get_matches(img_path string) []match {
	img_desc := lib.GetDescriptors(img_path)

	bf := gocv.NewBFMatcherWithParams(gocv.NormL1, false)
	defer bf.Close()

	var matches []match

	files, err := ioutil.ReadDir(KP_DIR)
	if err != nil {
		fmt.Println("Keypoints should be in data/kp")
		os.Exit(1)
	}

	i := 0
	for _, f := range files {
		if MAX_IMAGES > 0 && i >= MAX_IMAGES {
			break
		}

		kp_path := path.Join(KP_DIR, f.Name())
		desc := gocv.IMRead(kp_path, gocv.IMReadUnchanged)

		desc_matches := bf.KnnMatch(desc, img_desc, K)

		avg_dist := 0.0
		for _, m := range desc_matches {
			for _, j := range m {
				avg_dist += j.Distance
			}
		}
		avg_dist /= K
		matches = append(matches, match{kp_path, avg_dist})
		i++
	}

	sort.Slice(matches, func(i, j int) bool {
		return matches[i].avg_dist < matches[j].avg_dist
	})

	return matches
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Please provide an input image.")
		return
	}

	matches := get_matches(os.Args[1])

	fmt.Println("After Sorting Top 10 matches...")
	for _, m := range matches {
		fmt.Println(m.img, "\t", m.avg_dist)
	}
}
