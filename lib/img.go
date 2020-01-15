package lib

import (
	"gocv.io/x/gocv"
	"gocv.io/x/gocv/contrib"
)

func GetDescriptors(img_path string) (descs [][]float64) {
	sift := contrib.NewSIFT()
	defer sift.Close()

	mat := gocv.IMRead(img_path, gocv.IMReadGrayScale)
	defer mat.Close()

	// Extract descriptors
	if !mat.Empty() {
		mask := gocv.NewMat()
		key_pts, desc := sift.DetectAndCompute(mat, mask)

		for i, _ := range key_pts {
			var tmp []float64
			for j := 0; j < desc.Cols(); j++ {
				tmp = append(tmp, float64(desc.GetFloatAt(i, j)))
			}
			descs = append(descs, tmp)
		}

		mat.Close()
		desc.Close()
	}

	return descs
}
