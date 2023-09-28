package models

type RepoProcessor interface {
	ProcessRepo(Repo) error
}

type Repo struct {
	ID          int
	Name        string
	Description string
	UpdatedAt   string
	CloneURL    string
}
