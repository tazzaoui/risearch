package main

import (
	"fmt"
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

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	} else {
		target_ext := filepath.Ext(handler.Filename)
		ioutil.WriteFile("data/target"+target_ext, bytes, 0777)
	}

}

func main() {
	http.Handle("/", http.FileServer(http.Dir("html")))
	http.HandleFunc("/search", upload_file)
	fmt.Println("[+] Listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
