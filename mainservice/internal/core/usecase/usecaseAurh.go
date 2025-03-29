package usecase

import (
	"context"
	"fmt"
	"log/slog"
)

type UsecaseAuth struct {
	authClient         AuthClient
	notificationClient NotificationClient
}

func NewUsecaseAuth(authClient AuthClient, notificationClient NotificationClient) UsecaseAuth {
	return UsecaseAuth{
		authClient:         authClient,
		notificationClient: notificationClient,
	}
}

func (uc UsecaseAuth) Register(ctx context.Context, login string, hashPassword string) (string, error) {
	token, err := uc.authClient.Register(ctx, login, hashPassword)
	if err != nil {
		slog.Error("register error", "err", err)
		return "", err
	}

	err = uc.notificationClient.Publish(ctx, fmt.Sprintf(
		"User %s was register", login))

	if err != nil {
		slog.Error("send notification error", "err", err)

		return "", err
	}

	return token, nil
}

func (uc UsecaseAuth) Login(ctx context.Context, login string, hashPassword string) (string, error) {
	token, err := uc.authClient.Login(ctx, login, hashPassword)
	if err != nil {
		slog.Error("register error", "err", err)
		return "", err
	}

	err = uc.notificationClient.Publish(ctx, fmt.Sprintf(
		"User %s was login", login))

	if err != nil {
		slog.Error("send notification error", "err", err)

		return "", err
	}

	return token, nil
}
