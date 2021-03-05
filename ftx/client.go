package ftx

import (
	"net/http"
	"net/url"
)

// Client represents an API client.
type Client struct {
	Client     *http.Client
	API        string
	Secret     []byte
	Subaccount string
}

// NewClient creates a new API client.
func NewClient(api string, secret string, subaccount string) *Client {
	return &Client{Client: &http.Client{}, API: api, Secret: []byte(secret), Subaccount: url.PathEscape(subaccount)}
}
