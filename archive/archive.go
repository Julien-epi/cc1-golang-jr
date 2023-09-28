package archive

import (
	"archive/zip"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	csvDirPath := "../csvgenerate"
	csvFilePath := filepath.Join(csvDirPath, "repos.csv")
	zipFilePath := filepath.Join(csvDirPath, "repos.zip")

	err := CreateZip(csvFilePath, zipFilePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.ServeFile(w, r, zipFilePath)
}

func CreateZip(csvFilePath, zipFilePath string) error {
	zipFile, err := os.Create(zipFilePath)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	csvFile, err := os.Open(csvFilePath)
	if err != nil {
		return err
	}
	defer csvFile.Close()

	w, err := zipWriter.Create(filepath.Base(csvFilePath))
	if err != nil {
		return err
	}

	_, err = io.Copy(w, csvFile)
	return err
}
