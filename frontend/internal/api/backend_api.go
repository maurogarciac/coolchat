package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"frontend/config"
	"frontend/internal/domain"

	"go.uber.org/zap"
)

type BackendApi struct {
	client *http.Client
	cfg    *config.HTTPConfig
	lg     *zap.SugaredLogger
}

func NewBackendApi(client *http.Client, cfg *config.HTTPConfig, logger *zap.SugaredLogger) *BackendApi {
	return &BackendApi{client: client, cfg: cfg, lg: logger}
}

func (p *BackendApi) PostLogin(ctx context.Context, input domain.User) (domain.LoginResult, error) {

	ctx, cancel := context.WithTimeout(ctx, p.cfg.Timeout)
	defer cancel()

	url := p.cfg.BackendAPIURL + "/login"
	p.lg.Debugf("backend-api POST login request to url=%s", url)

	body := toBackendApiLoginRequest(input)

	reqBody, err := json.Marshal(body)
	if err != nil {
		return domain.LoginResult{}, fmt.Errorf("backend-api POST login error: %w", err)
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody)) // Create POST request
	if err != nil {
		return domain.LoginResult{}, fmt.Errorf("backend-api POST login error: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req = req.WithContext(ctx)

	resp, err := p.client.Do(req) // Perform request
	if err != nil {
		return domain.LoginResult{}, fmt.Errorf("backend-api POST login error: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return domain.LoginResult{}, fmt.Errorf("backend-api POST login req failed with status=%s", resp.Status)
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return domain.LoginResult{}, fmt.Errorf("backend-api POST login read resp error: %w", err)
	}

	var loginResp PostLoginResponse
	err = json.Unmarshal(respBody, &loginResp) // Convert values to local type
	if err != nil {
		return domain.LoginResult{}, fmt.Errorf("backend-api POST login unmarshal resp error: %w", err)
	}
	p.lg.Info(loginResp)
	return loginResp.PostLoginResult(), nil
}

func (p *BackendApi) GetMessageHistory(ctx context.Context) (domain.MessageHistoryResult, error) {
	ctx, cancel := context.WithTimeout(ctx, p.cfg.Timeout)
	defer cancel()

	url := p.cfg.BackendAPIURL + "/messages"
	p.lg.Debugf("backend-api GET messages request to url=%s", url)

	req, err := http.NewRequest("GET", url, nil) // Create POST request
	if err != nil {
		return domain.MessageHistoryResult{}, fmt.Errorf("backend-api GET messages error: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req = req.WithContext(ctx)

	resp, err := p.client.Do(req) // Perform request
	if err != nil {
		return domain.MessageHistoryResult{}, fmt.Errorf("backend-api GET messages error: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return domain.MessageHistoryResult{}, fmt.Errorf("backend-api GET messages req failed with status=%s", resp.Status)
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return domain.MessageHistoryResult{}, fmt.Errorf("backend-api GET messages read resp error: %w", err)
	}

	var valueResp MessageHistoryResponse
	err = json.Unmarshal(respBody, &valueResp) // Convert values to local type
	if err != nil {
		return domain.MessageHistoryResult{}, fmt.Errorf("backend-api GET messages unmarshal resp error: %w", err)
	}

	return valueResp.MessageHistoryResult(), nil
}
