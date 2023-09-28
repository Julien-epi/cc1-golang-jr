package github

import (
	"cc1-jr/models"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
	"sync"
	"time"
)

type GitHubError struct {
	Message string `json:"message"`
}

func GenerateReposCSV(processor models.RepoProcessor) error {
	orgURL, urlSet := os.LookupEnv("GITHUB_ORG_URL")
	if !urlSet {
		log.Fatal("GITHUB_ORG_URL environment variable is not set")
	}
	url := fmt.Sprintf("%s/repos", orgURL)

	token, tokenSet := os.LookupEnv("PERSONNAL_TOKEN_GITHUB")
	if !tokenSet {
		log.Fatal("PERSONNAL_TOKEN_GITHUB environment variable is not set")
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", "token "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		var errResp GitHubError
		if err := json.Unmarshal(body, &errResp); err != nil {
			return err
		}
		return fmt.Errorf("GitHub Error: %s", errResp.Message)
	}

	var repos []models.Repo
	if err := json.Unmarshal(body, &repos); err != nil {
		return err
	}

	sortReposByDate(repos)

	var wg sync.WaitGroup
	for _, repo := range repos {
		wg.Add(1)
		go func(repo models.Repo) {
			defer wg.Done()
			if err := processor.ProcessRepo(repo); err != nil {
				log.Println("Error processing repo: ", err)
			}
		}(repo)
	}
	wg.Wait()

	return nil
}

func sortReposByDate(repos []models.Repo) {
	sort.Slice(repos, func(i, j int) bool {
		dateI, _ := time.Parse(time.RFC3339, repos[i].UpdatedAt)
		dateJ, _ := time.Parse(time.RFC3339, repos[j].UpdatedAt)
		return dateI.After(dateJ)
	})
}
