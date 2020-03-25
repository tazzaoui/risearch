package config

import (
	"os"
	"path"
)

// Path to reference bank
func ImgDir() string {
	return path.Join(os.Getenv("GOPATH"), "src/github.com/tazzaoui/risearch/data/img")
}

// Path to their extracted keypoints
func KpDir() string {
	return path.Join(os.Getenv("GOPATH"), "src/github.com/tazzaoui/risearch/data/kp")
}

// Maximum number of images index
func MaxImages() int {
	return 100
}

// Count of best matches found per each query descriptor. Used in knnMatch()
// See: https://docs.opencv.org/master/db/d39/classcv_1_1DescriptorMatcher.html#aa880f9353cdf185ccf3013e08210483a
func K() int {
	return 4
}
