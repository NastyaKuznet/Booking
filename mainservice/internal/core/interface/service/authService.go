package service

import (
	"context"
	"crypto/sha1"
	"errors"
	"fmt"
	"log/slog"
	"mainservice/internal/core/entity"
	"mainservice/internal/repository/datasources/grpc"
	"time"

	"github.com/golang-jwt/jwt"
)

type TokenClaims struct {
	jwt.StandardClaims
	Login string `json:"login"`
}

type AuthService struct {
	repo grpc.AuthRepository
}

func NewAuthService(repo grpc.AuthRepository) AuthService {
	return AuthService{repo: repo}
}

func (service AuthService) Register(ctx context.Context, login string,
	password string) (string, error) {

	hash := generatePassword(password)

	userName, err := service.repo.Register(ctx, login, hash)

	if err != nil {
		slog.Error(err.Error())
		return "", errors.New("user not created")
	}

	return generateToken(userName)
}

func generatePassword(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(entity.Salt)))
}

func generateToken(login string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(entity.TokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		Login: login,
	})
	return token.SignedString([]byte(entity.SignInKey))
}
