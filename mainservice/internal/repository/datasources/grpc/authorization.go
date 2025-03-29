package grpc

import (
	"context"
	"fmt"
)

type AuthRepository struct {
	client AuthClient
}

func NewAuthRepository(client AuthClient) *AuthRepository {
	return &AuthRepository{client: client}
}

func (b *AuthRepository) Login(ctx context.Context, login string, password string) (string, error) {
	token, err := b.client.Login(ctx, login, password)

	if err != nil {
		return "", fmt.Errorf("%w: %w", ErrInternal, err)
	}
	return token, err
}

func (b *AuthRepository) Register(ctx context.Context, login string, password string) (string, error) {
	token, err := b.client.Register(ctx, login, password)

	if err != nil {
		return "", fmt.Errorf("%w: %w", ErrInternal, err)
	}
	return token, err
}
