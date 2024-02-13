package hstorage_sdk

import (
	"fmt"
	"io"
	"net/http"
)

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	req.Header.Add("x-eu-api-key", c.APIKey)
	req.Header.Add("x-eu-email", c.Email)

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("HTTP error: %s", resp.Status)
	}

	return body, nil
}
