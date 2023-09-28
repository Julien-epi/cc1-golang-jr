package main

import (
	"cc1-jr/archive"
	"cc1-jr/csv"
	"cc1-jr/github"
	"fmt"
	"log"
	"net/http"
)

func main() {
	processor := &csv.CSVProcessor{}
	err := github.GenerateReposCSV(processor)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/archive", archive.Handler)
	http.HandleFunc("/download", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "csvgenerate/repos.zip")
	})

	fmt.Println("Server is running on port 8081...")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
