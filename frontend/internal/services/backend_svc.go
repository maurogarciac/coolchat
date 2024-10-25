package services

import (
	"context"

	"frontend/config"
	"frontend/internal/api"
	"frontend/internal/domain"
)

type BackendService struct {
	cfg     *config.AppConfig
	backend *api.BackendApi
}

func NewBackendService(
	config *config.AppConfig,
	backendApi *api.BackendApi,
) *BackendService {
	return &BackendService{
		cfg:     config,
		backend: backendApi,
	}
}

// Issue a POST request to the Backend to retrieve JWT Access and Refresh tokens
func (b *BackendService) PostLogin(ctx context.Context, input domain.User) (domain.LoginResult, error) {

	res, err := b.backend.PostLogin(ctx, input)
	if err != nil {
		return domain.LoginResult{}, err
	}

	return res, nil
}

// Issue a POST request to the Backend to retrieve an updated Access Token
func (b *BackendService) PostRefresh(ctx context.Context, input domain.RefreshToken) (domain.RefreshResult, error) {

	res, err := b.backend.PostRefresh(ctx, input)
	if err != nil {
		return domain.RefreshResult{}, err
	}

	return res, nil
}

// Issue a GET request to the Backend to retrieve chat message history
func (b *BackendService) GetMessageHistory(ctx context.Context) (domain.MessageHistoryResult, error) {

	messages, err := b.backend.GetMessageHistory(ctx)

	if err != nil {
		return domain.MessageHistoryResult{}, err
	}

	return messages, nil
}
