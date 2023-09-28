package csv

import (
	"archive/zip"
	"cc1-jr/models"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type CSVProcessor struct{}

func (c *CSVProcessor) ProcessRepo(repo models.Repo) error {
	csvFilePath, err := writeToCSV(repo)
	if err != nil {
		return err
	}

	err = archiveRepos(csvFilePath)
	if err != nil {
		return err
	}
	return nil
}

func writeToCSV(repo models.Repo) (string, error) {
	csvDirPath := "../csvgenerate"
	csvFilePath := filepath.Join(csvDirPath, "repos.csv")
	if _, err := os.Stat(csvDirPath); os.IsNotExist(err) {
		err = os.MkdirAll(csvDirPath, 0755)
		if err != nil {
			return "", err
		}
	}

	file, err := os.OpenFile(csvFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return "", err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	info, err := file.Stat()
	if err != nil {
		return "", err
	}
	if info.Size() == 0 {
		err = writer.Write([]string{"ID", "Name", "Description", "URL"})
		if err != nil {
			return "", err
		}
	}

	err = writer.Write([]string{
		fmt.Sprintf("%d", repo.ID),
		repo.Name,
		repo.Description,
		repo.CloneURL,
	})
	if err != nil {
		return "", err
	}

	return csvFilePath, nil
}

func archiveRepos(csvFilePath string) error {
	dir := filepath.Dir(csvFilePath)
	archivePath := filepath.Join(dir, "repos.zip")
	archiveFile, err := os.Create(archivePath)
	if err != nil {
		return err
	}
	defer archiveFile.Close()

	zipWriter := zip.NewWriter(archiveFile)
	defer zipWriter.Close()

	file, err := os.Open(csvFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	w, err := zipWriter.Create(filepath.Base(csvFilePath))
	if err != nil {
		return err
	}

	_, err = io.Copy(w, file)
	return err
}
