package main

import (
    "fmt"
    "os"
    "io/ioutil"
)

func main() {
    image_dir := "img"

    images, err := ioutil.ReadDir(image_dir)

    if err != nil {
        fmt.Fprintf(os.Stderr, "ERROR %v\n", err)
        os.Exit(1)
    }

    for _, img := range images {
        fmt.Println(img.Name())
    }
}
