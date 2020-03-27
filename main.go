package main

import (
	"fmt"
	"github.com/tazzaoui/risearch/lib"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Please provide an input image.")
		return
	}

	matches := lib.GetMatches(os.Args[1])

	fmt.Println("After Sorting Top 10 matches...")
	for _, m := range matches {
		fmt.Println(m.Img, "\t", m.Sim)
	}
}
