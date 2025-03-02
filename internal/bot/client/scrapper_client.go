package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/RobertGabdullin/GoTelegramBot/internal/transport"
	"log"
	"net/http"
)

type ScrapperClient struct {
	BaseUrl    string
	HttpClient *http.Client
}

func NewScrapperClient(baseUrl string) *ScrapperClient {
	return &ScrapperClient{
		BaseUrl: baseUrl,
	}
}

func (c *ScrapperClient) PostTgChat(chatId int64) (string, error) {
	req, err := http.NewRequest(http.MethodPost, c.BaseUrl+"/tg-chat/"+fmt.Sprint(chatId), nil)
	if err != nil {
		return "", fmt.Errorf("failed to register. Try later")
	}
	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("internal server error. Try later")
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		decoder := json.NewDecoder(resp.Body)
		var errorResponse transport.APIErrorResponse
		err = decoder.Decode(&errorResponse)
		if err != nil {
			return "", fmt.Errorf("failed to register. Try later")
		}
		return errorResponse.Description, nil
	}

	return "You have successfully registered", nil

}

func (c *ScrapperClient) DeleteTgChat() {

}

func (c *ScrapperClient) GetLinks(chatId int64) ([]transport.LinkResponse, error) {
	req, err := http.NewRequest(http.MethodGet, c.BaseUrl+"/links", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get links. Try later")
	}
	req.Header.Set("Tg-Chat-Id", fmt.Sprint(chatId))

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get links. Try later")
	}

	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)

	if resp.StatusCode != http.StatusOK {
		var errorResponse transport.APIErrorResponse
		err = decoder.Decode(&errorResponse)
		if err != nil {
			return nil, fmt.Errorf("Internal server error. Try later")
		}
		return nil, errors.New(errorResponse.Description)
	}

	var listLinksResponse transport.ListLinksResponse
	err = decoder.Decode(&listLinksResponse)
	if err != nil {
		return nil, fmt.Errorf("internal server error. Try later")
	}

	return listLinksResponse.Links, nil

}

func (c *ScrapperClient) PostLinks(chatId int64, link string) (string, error) {
	addLink := transport.AddLinkRequest{
		Link: link,
	}
	body, err := json.Marshal(addLink)
	req, err := http.NewRequest(http.MethodPost, c.BaseUrl+"/links", bytes.NewReader(body))
	if err != nil {
		return "", fmt.Errorf("failed to add link. Try later")
	}

	req.Header.Set("Tg-Chat-Id", fmt.Sprint(chatId))
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("internal server error. Try later")
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		decoder := json.NewDecoder(resp.Body)
		var errorResponse transport.APIErrorResponse
		err = decoder.Decode(&errorResponse)
		if err != nil {
			return "", fmt.Errorf("failed to add link. Try later")
		}
		return "", errors.New(errorResponse.Description)
	}

	return "You have successfully started tracking the link", nil

}

func (c *ScrapperClient) DeleteLinks(chatId int64, link string) (string, error) {
	removeLink := transport.RemoveLinkRequest{
		Link: link,
	}

	body, err := json.Marshal(removeLink)
	if err != nil {
		return "", fmt.Errorf("failed to remove link. Try later")
	}

	req, err := http.NewRequest(http.MethodDelete, c.BaseUrl+"/links", bytes.NewReader(body))
	if err != nil {
		return "", fmt.Errorf("failed to remove link. Try later")
	}

	req.Header.Set("Tg-Chat-Id", fmt.Sprint(chatId))
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.HttpClient.Do(req)
	if err != nil {
		log.Printf("failed to remove link. Try later")
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		decoder := json.NewDecoder(resp.Body)
		var errorResponse transport.APIErrorResponse
		err = decoder.Decode(&errorResponse)
		if err != nil {
			return "", fmt.Errorf("failed to remove link. Try later")
		}
		return "", errors.New(errorResponse.Description)
	}

	return "You have successfully removed the link", nil
}
