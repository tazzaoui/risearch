package main

import (
    "fmt"
    "os"
    "io/ioutil"
)

func main() {
    image_dir := "data/img"

    images, err := ioutil.ReadDir(image_dir)

    if err != nil {
        fmt.Println("Please download the image data set by running get_data.sh")
        os.Exit(1)
    }

    for _, img := range images {
        fmt.Println(img.Name())
    }
}
