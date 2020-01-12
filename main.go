package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) <= 1 {
		fmt.Println("Please provide an input image.")
		return
	}

	img_path := os.Args[1]

	fmt.Println(img_path)
}
