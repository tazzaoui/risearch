package main

import (
	"fmt"
	"github.com/tazzaoui/risearch/lib"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

func upload_file(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[+] get_file_path")
	r.ParseMultipartForm(32 << 20)

	file, handler, err := r.FormFile("uploadfile")
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(handler.Filename)

	// Copy the query image data/target.ext
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	target_path := "data/target" + filepath.Ext(handler.Filename)
	ioutil.WriteFile(target_path, bytes, 0777)

	matches := lib.GetMatches(target_path)
	fmt.Println(matches)
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("html")))
	http.HandleFunc("/search", upload_file)
	fmt.Println("[+] Listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
