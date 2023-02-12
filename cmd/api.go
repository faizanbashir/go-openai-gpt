package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) Engines(ctx context.Context) (*EnginesResponse, error) {
	response := c.callAPI(ctx, "GET", "engines", nil)
	if response.StatusCode == http.StatusOK {
		enginesReponse := new(EnginesResponse)
		data := response.Response
		json.Unmarshal([]byte(data), &enginesReponse)
		return enginesReponse, nil
	}
	return nil, fmt.Errorf(string(response.Response))
}

func (c *Client) Engine(ctx context.Context, engine string) (*EngineObject, error) {
	response := c.callAPI(ctx, "GET", fmt.Sprintf("engines/%s", engine), nil)
	if response.StatusCode == http.StatusOK {
		enginesReponse := new(EngineObject)
		data := response.Response
		json.Unmarshal([]byte(data), &enginesReponse)
		return enginesReponse, nil
	}
	return nil, fmt.Errorf(string(response.Response))
}
