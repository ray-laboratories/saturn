package remote

import (
	bytes "bytes"
	"context"
	"encoding/json"
	"github.com/ray-laboratories/saturn/application"
	"github.com/ray-laboratories/saturn/types"
	"io"
	"net/http"
	"strings"
)

var _ application.AuthService = &AuthService{}

type AuthService struct {
	BaseURL string
	Client  *http.Client
}

func (a *AuthService) Login(ctx context.Context, username, password string) (string, error) {
	resp, err := a.request("/login", "POST", map[string]any{username: username, password: password})
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	// Unmarshal body
	type tokenResponse struct {
		Token string `json:"token"`
	}
	var token tokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&token); err != nil {
		return "", err
	}
	return token.Token, nil
}

func (a *AuthService) Logout(ctx context.Context, token string) error {
	resp, err := a.request("/logout", "POST", nil, "Authorization", "Bearer "+token)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func (a *AuthService) Register(ctx context.Context, username, password string) (string, error) {
	resp, err := a.request("/register", "POST", map[string]any{username: username, password: password})
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	// Unmarshal body
	type tokenResponse struct {
		Token string `json:"token"`
	}
	var token tokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&token); err != nil {
		return "", err
	}
	return token.Token, nil
}

func (a *AuthService) Validate(ctx context.Context, token string) (*types.Accessor, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AuthService) request(url, method string, body any, headers ...string) (*http.Response, error) {
	var reqBody io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		reqBody = bytes.NewReader(b)
	}
	req, err := http.NewRequest(method, a.BaseURL+url, reqBody)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	for i := 0; i+1 < len(headers); i += 2 {
		req.Header.Set(headers[i], headers[i+1])
	}
	return a.Client.Do(req)
}

func NewAuthService(baseURL string, client *http.Client) *AuthService {
	return &AuthService{
		BaseURL: strings.TrimRight(baseURL, "/"),
		Client:  client,
	}
}
