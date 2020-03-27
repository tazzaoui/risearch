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
	Img string  // Path to image being matched against
	Sim float64 // Similarity metric
}

// Ratio used to filter keypts
const RATIO = 0.7

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

		var valid_pts []gocv.DMatch

		for _, m := range desc_matches {
			// My dude David Lowe's ratio test
			if m[0].Distance < RATIO*m[1].Distance {
				valid_pts = append(valid_pts, m[0])
			}
		}

		similarity := float64(len(valid_pts)) / float64(len(desc_matches))

		img_path := config.ImgDir() + "/" + f.Name()[:len(f.Name())-5]
		matches = append(matches, Match{img_path, similarity})
		i++
	}

	sort.Slice(matches, func(i, j int) bool {
		return matches[i].Sim > matches[j].Sim
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
