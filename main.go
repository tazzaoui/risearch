package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/tazzaoui/risearch/lib"
)

func read_descs(max int) [][][]float64 {
	kp_dir := "data/kp"
	files, err := ioutil.ReadDir(kp_dir)

	if err != nil {
		fmt.Println("Keypoints should be in data/kp")
		os.Exit(1)
	}

	var descs [][][]float64

	i := 0
	for _, f := range files {
		if max > 0 && i >= max {
			break
		}
		kp_path := path.Join(kp_dir, f.Name())
		var tmp [][]float64
		file, _ := ioutil.ReadFile(kp_path)
		_ = json.Unmarshal(file, &tmp)
		descs = append(descs, tmp)
		i++
	}
	return descs
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Please provide an input image.")
		return
	}

	desc_db := read_descs(100)
	fmt.Println(desc_db)

	img_desc := lib.GetDescriptors(os.Args[1])
	fmt.Println(img_desc)
}
