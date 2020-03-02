package main

import (
	"fmt"
	"github.com/tazzaoui/risearch/lib"
	"html/template"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

type data struct {
	Query   string      // Path of the query image
	Matches []lib.Match // Matched images
}

func UploadFile(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)

	file, handler, err := r.FormFile("uploadfile")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fmt.Println("[+] Querying Image: " + handler.Filename + "...")

	// Copy the query image to data/target.ext
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	target_path := "data/target" + filepath.Ext(handler.Filename)
	ioutil.WriteFile(target_path, bytes, 0777)

	template := template.Must(template.ParseFiles("templates/results.html"))

	matches := lib.GetMatches(target_path)
	fmt.Printf("[+] Matches: %d Closest Match: %.2f\n", len(matches), matches[0].AvgDist)

	err = template.ExecuteTemplate(w, "results.html", data{target_path, matches})
	if err != nil {
		fmt.Println(err)
		return
	}

}

func main() {
	http.Handle("/", http.FileServer(http.Dir("templates")))
	http.HandleFunc("/search", UploadFile)

	// Serve static image files (for displaying results)
	img_fs := http.FileServer(http.Dir("data"))
	http.Handle("/data/", http.StripPrefix("/data/", img_fs))

	fmt.Println("[+] Listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
