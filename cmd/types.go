package cli

import (
	"context"
	"net/http"
)

type Client struct {
	Endpoint        string
	HTTPClient      *http.Client
	SecretAccessKey string
	AIEngine        string
	idOrg           string
}

type APIResponse struct {
	StatusCode int
	Response   []uint8
}

// EngineObject contained in an engine reponse
type EngineObject struct {
	ID     string `json:"id"`
	Object string `json:"object"`
	Owner  string `json:"owner"`
	Ready  bool   `json:"ready"`
}

// EnginesResponse is returned from the Engines API
type EnginesResponse struct {
	Data   []EngineObject `json:"data"`
	Object string         `json:"object"`
}

// A Client is an API client to communicate with the OpenAI gpt-3 APIs
type ClientInterface interface {
	// Engines lists the currently available engines, and provides basic information about each
	// option such as the owner and availability.
	Engines(ctx context.Context) (*EnginesResponse, error)

	// Engine retrieves an engine instance, providing basic information about the engine such
	// as the owner and availability.
	Engine(ctx context.Context, engine string) (*EngineObject, error)
}
