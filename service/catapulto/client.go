package catapulto

import (
	"net/http"
)

type Client struct {
	baseURL    string
	xToken     string
	HTTPClient *http.Client
}

func NewClient(
	baseURL string,
	xToken string,
) *Client {
	return &Client{
		baseURL: baseURL,
		xToken:  xToken,
	}
}
