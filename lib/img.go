package lib

import (
	"gocv.io/x/gocv"
	"gocv.io/x/gocv/contrib"
)

func GetDescriptors(img_path string) gocv.Mat {
	sift := contrib.NewSIFT()
	defer sift.Close()

	mat := gocv.IMRead(img_path, gocv.IMReadGrayScale)
	defer mat.Close()

	mask := gocv.NewMat()
	_, desc := sift.DetectAndCompute(mat, mask)
	mask.Close()

	return desc // IT IS THE CALLER'S RESPONSIBILITY TO CLOSE THIS
}
