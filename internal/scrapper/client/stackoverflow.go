package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type StackOverflowClient struct {
	ApiUrl   string
	Question string
}

type stackoverflowDTO struct {
	Items []struct {
		LastActivity int64 `json:"last_activity_date"`
	} `json:"items"`
}

func NewStackOverflowClient(questionUrl string) *StackOverflowClient {
	apiUrl := "https://api.stackexchange.com/2.3"
	temp := strings.Split(questionUrl, "://")[1]
	question := strings.Split(temp, "/")[2]
	return &StackOverflowClient{
		ApiUrl:   apiUrl,
		Question: question,
	}
}

func (client StackOverflowClient) Scrape() (time.Time, error) {

	httpClient := &http.Client{Timeout: 10 * time.Second}

	req, err := http.NewRequest("GET",
		client.ApiUrl+"/questions/"+client.Question+"?site=stackoverflow", nil)
	if err != nil {
		return time.Time{}, fmt.Errorf("create request failed: %w", err)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return time.Time{}, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return time.Time{}, fmt.Errorf("stackoverflow return non 200 status code: %d", resp.StatusCode)
	}

	var result stackoverflowDTO
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return time.Time{}, fmt.Errorf("decoding response failed: %w", err)
	}

	t := time.Unix(result.Items[0].LastActivity, 0)
	t = t.Truncate(time.Second)

	return t, nil
}
