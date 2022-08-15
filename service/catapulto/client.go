package catapulto

import (
	"bytes"
	"encoding/json"
	"fmt"
	"golang/service/catapulto/api"
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
		HTTPClient: &http.Client{},
		baseURL:    baseURL,
		xToken:     xToken,
	}
}

func (client *Client) GetAuthToken(authRequest *api.AuthRequest) (*api.AuthResponse, error) {
	body := new(bytes.Buffer)
	err := json.NewEncoder(body).Encode(authRequest)
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf(
		"%s/users/api-token-auth/",
		client.baseURL,
	)
	req, err := http.NewRequest(
		http.MethodPost,
		url,
		body,
	)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}
	res, err := client.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	var authResponse *api.AuthResponse
	err = json.NewDecoder(res.Body).Decode(&authResponse)
	if err != nil {
		return nil, err
	}
	err = res.Body.Close()
	if err != nil {
		return nil, err
	}

	return authResponse, nil
}

func (client *Client) GetLocalityList(localityRequest *api.LocalityRequest) ([]api.LocalityResponse, error) {
	url := fmt.Sprintf(
		"%s/geo/locality/search/?term=%s&ISO=%s&limit=%d",
		client.baseURL,
		localityRequest.Term,
		localityRequest.ISO,
		localityRequest.Limit,
	)
	req, err := http.NewRequest(
		http.MethodGet,
		url,
		nil,
	)
	if err != nil {
		return nil, err
	}
	res, err := client.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	var localityList []api.LocalityResponse
	err = json.NewDecoder(res.Body).Decode(&localityList)
	if err != nil {
		return nil, err
	}
	err = res.Body.Close()
	if err != nil {
		return nil, err
	}

	return localityList, nil
}
