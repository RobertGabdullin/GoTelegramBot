package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type GitHubClient struct {
	ApiUrl   string
	UserRepo string
}

type gitHubDTO struct {
	LastUpdate time.Time `json:"timestamp"`
}

func NewGitHubClient(repoUrl string) *GitHubClient {
	userRepo := strings.Split(repoUrl, ".com")[1]
	return &GitHubClient{
		ApiUrl:   "https://api.github.com",
		UserRepo: userRepo,
	}
}

func (client GitHubClient) Scrape() (time.Time, error) {
	httpClient := http.Client{Timeout: time.Second * 10}
	req, err := http.NewRequest(http.MethodGet,
		client.ApiUrl+"/repos"+client.UserRepo+"/activity", nil)
	if err != nil {
		return time.Time{}, err
	}

	req.Header.Add("Accept", "application/vnd.github.v3+json")
	resp, err := httpClient.Do(req)
	if err != nil {
		return time.Time{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return time.Time{}, errors.New("GitHub returned non-200 status code")
	}

	var result []gitHubDTO
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return time.Time{}, fmt.Errorf("json decode failed: %w", err)
	}

	resultTime := result[0].LastUpdate
	resultTime = resultTime.Truncate(time.Second)

	return resultTime, nil
}
