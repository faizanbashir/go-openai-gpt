package cli

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const (
	OpenAIEndpoint = "https://api.openai.com/v1"
	DefaultTimeout = 60
	ChatGPTEngine  = "text-chat-davinci-002-20230126"
)

func transport() *http.Client {
	return &http.Client{
		Timeout: time.Duration(DefaultTimeout * time.Second),
	}
}

func GetClient(secretKey, org string) *Client {
	return &Client{
		Endpoint:        OpenAIEndpoint,
		HTTPClient:      transport(),
		SecretAccessKey: secretKey,
		AIEngine:        ChatGPTEngine,
		idOrg:           org,
	}
}

func (c *Client) callAPI(ctx context.Context, method, path string, payload interface{}) APIResponse {
	body, err := jsonBodyParser(payload)
	if err != nil {
		return APIResponse{
			StatusCode: 500,
			Response:   []uint8("{}"),
		}
	}
	url := fmt.Sprintf("%s/%s", c.Endpoint, path)
	fmt.Printf("Calling API: %s\n", url)
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return APIResponse{
			StatusCode: 500,
			Response:   []uint8("{}"),
		}
	}
	if len(c.idOrg) > 0 {
		req.Header.Set("OpenAI-Organization", c.idOrg)
	}
	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.SecretAccessKey))
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		if os.IsTimeout(err) {
			return APIResponse{
				StatusCode: 500,
				Response:   []uint8("{}"),
			}
		}
	}
	defer resp.Body.Close()
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return APIResponse{
			StatusCode: resp.StatusCode,
			Response:   content,
		}
	}
	return APIResponse{
		StatusCode: resp.StatusCode,
		Response:   content,
	}
}

func jsonBodyParser(body interface{}) (io.Reader, error) {
	if body == nil {
		return bytes.NewBuffer(nil), nil
	}
	raw, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("failed encoding json: %w", err)
	}
	return bytes.NewBuffer(raw), nil
}
