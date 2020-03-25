package lib

import (
	"fmt"
	"github.com/tazzaoui/risearch/config"
	"gocv.io/x/gocv"
	"gocv.io/x/gocv/contrib"
	"io/ioutil"
	"os"
	"path"
	"sort"
)

type Match struct {
	Img     string  // Path to image being matched against
	AvgDist float64 // Average Euclidean Distance between matches
}

func GetMatches(img_path string) []Match {
	img_desc := GetDescriptors(img_path)

	bf := gocv.NewBFMatcherWithParams(gocv.NormL1, false)
	defer bf.Close()

	var matches []Match

	files, err := ioutil.ReadDir(config.KpDir())
	if err != nil {
		fmt.Println("Keypoints should be in ", config.KpDir())
		os.Exit(1)
	}

	i := 0
	for _, f := range files {
		if config.MaxImages() > 0 && i >= config.MaxImages() {
			break
		}

		kp_path := path.Join(config.KpDir(), f.Name())
		desc := gocv.IMRead(kp_path, gocv.IMReadUnchanged)

		desc_matches := bf.KnnMatch(desc, img_desc, config.K())

		avg_dist := 0.0
		for _, m := range desc_matches {
			for _, j := range m {
				avg_dist += j.Distance
			}
		}
		avg_dist /= float64(config.K())
		img_path := config.ImgDir() + "/" + f.Name()[:len(f.Name())-5]
		matches = append(matches, Match{img_path, avg_dist})
		i++
	}

	sort.Slice(matches, func(i, j int) bool {
		return matches[i].AvgDist < matches[j].AvgDist
	})

	return matches
}

func GetDescriptors(img_path string) gocv.Mat {
	sift := contrib.NewSIFT()
	defer sift.Close()

	mat := gocv.IMRead(img_path, gocv.IMReadGrayScale)
	defer mat.Close()

	mask := gocv.NewMat()
	_, desc := sift.DetectAndCompute(mat, mask)
	mask.Close()

	return desc
}
