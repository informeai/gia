package dto

type Author struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	CommitCount int    `json:"commit_count"`
}
