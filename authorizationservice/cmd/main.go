package main

import (
	"context"
	"crypto/sha1"
	"errors"
	"fmt"
	"log"
	"net"
	"time"

	"authorizationservice/internal/middleware"
	"authorizationservice/order"

	"github.com/golang-jwt/jwt"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
)

const (
	connection = "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s"
	salt       = "your-salt-here"
	signingKey = "your-signing-key-here"
	tokenTTL   = 12 * time.Hour
)

type AuthServer struct {
	order.UnimplementedAuthServiceServer
	db *pgxpool.Pool
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	conString := fmt.Sprintf(connection, "postgres", "5432", "postgres", "postgres", "postgres", "disable")
	conn, err := pgxpool.New(context.Background(), conString)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.AuthInterceptor),
	)
	order.RegisterAuthServiceServer(s, &AuthServer{db: conn})
	log.Println("Starting auth service...")
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *AuthServer) Register(ctx context.Context, req *order.RegisterRequest) (*order.AuthResponse, error) {
	hash := generatePasswordHash(req.Password)

	var userID int
	err := s.db.QueryRow(ctx,
		"INSERT INTO users (login, password_hash) VALUES ($1, $2) RETURNING id",
		req.Login, hash).Scan(&userID)

	if err != nil {
		return &order.AuthResponse{Error: "registration failed"}, err
	}

	token, err := generateToken(req.Login)
	if err != nil {
		return &order.AuthResponse{Error: "failed to generate token"}, err
	}

	return &order.AuthResponse{Token: token}, nil
}

func (s *AuthServer) Login(ctx context.Context, req *order.LoginRequest) (*order.AuthResponse, error) {
	var (
		dbHash string
		login  string
	)

	err := s.db.QueryRow(ctx,
		"SELECT login, password_hash FROM users WHERE login = $1",
		req.Login).Scan(&login, &dbHash)

	if err != nil {
		return &order.AuthResponse{Error: "user not found"}, err
	}

	if generatePasswordHash(req.Password) != dbHash {
		return &order.AuthResponse{Error: "invalid credentials"}, errors.New("invalid password")
	}

	token, err := generateToken(login)
	if err != nil {
		return &order.AuthResponse{Error: "failed to generate token"}, err
	}

	return &order.AuthResponse{Token: token}, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func generateToken(login string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(tokenTTL).Unix(),
		IssuedAt:  time.Now().Unix(),
		Subject:   login,
	})

	return token.SignedString([]byte(signingKey))
}
