package hstorage_sdk

import (
	"errors"
	"net/http"
)

type Client struct {
	APIKey     string
	Email      string
	BaseURL    string
	HttpClient *http.Client
}

func New(apiKey, email string, httpClient *http.Client) (*Client, error) {
	if apiKey == "" || email == "" {
		return nil, errors.New("apiKey and email are required")
	}

	if httpClient == nil {
		httpClient = &http.Client{}
	}

	return &Client{
		APIKey:     apiKey,
		Email:      email,
		BaseURL:    "https://api.hstorage.io",
		HttpClient: httpClient,
	}, nil
}
