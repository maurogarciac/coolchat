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

func (b *BackendService) PostLogin(ctx context.Context, input domain.User) (domain.LoginResult, error) {

	res, err := b.backend.PostLogin(ctx, input)
	if err != nil {
		return domain.LoginResult{}, err
	}

	return res, nil
}

func (b *BackendService) GetMessageHistory(ctx context.Context) (domain.MessageHistoryResult, error) {

	messages, err := b.backend.GetMessageHistory(ctx)
	if err != nil {
		return domain.MessageHistoryResult{}, err
	}

	return messages, nil
}
