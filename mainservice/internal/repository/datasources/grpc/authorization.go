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

func (b *AuthRepository) GetUser(ctx context.Context, login string, hashPassword string) (string, error) {
	login, err := b.client.GetUser(ctx, login, hashPassword)

	if err != nil {
		return "", fmt.Errorf("%w: %w", ErrInternal, err)
	}
	return login, err
}

func (b *AuthRepository) Register(ctx context.Context, login string, hashPassword string) (string, error) {
	login, err := b.client.Register(ctx, login, hashPassword)

	if err != nil {
		return "", fmt.Errorf("%w: %w", ErrInternal, err)
	}
	return login, err
}
