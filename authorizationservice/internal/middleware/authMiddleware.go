package middleware

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	signingKey = "your-signing-key-here"
)

type contextKey string

const (
	userClaimsKey contextKey = "user_claims"
)

func AuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if info.FullMethod == "/order.AuthService/Register" ||
		info.FullMethod == "/order.AuthService/Login" ||
		info.FullMethod == "/order.AuthService/ValidateToken" {
		return handler(ctx, req)
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "missing metadata")
	}

	authHeaders := md.Get("authorization")
	if len(authHeaders) == 0 {
		return nil, status.Error(codes.Unauthenticated, "authorization token is required")
	}

	tokenString := strings.TrimPrefix(authHeaders[0], "Bearer ")

	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(signingKey), nil
	})

	if err != nil {
		log.Printf("Token validation error: %v", err)
		return nil, status.Error(codes.Unauthenticated, "invalid token")
	}

	if !token.Valid {
		return nil, status.Error(codes.Unauthenticated, "invalid token")
	}

	if claims, ok := token.Claims.(*jwt.StandardClaims); ok {
		ctx = context.WithValue(ctx, userClaimsKey, claims)
	} else {
		return nil, status.Error(codes.Unauthenticated, "invalid token claims")
	}

	return handler(ctx, req)
}

func GetLoginFromContext(ctx context.Context) (string, error) {
	claims, ok := ctx.Value(userClaimsKey).(*jwt.StandardClaims)
	if !ok {
		return "", fmt.Errorf("could not get claims from context")
	}

	if time.Now().Unix() > claims.ExpiresAt {
		return "", fmt.Errorf("token expired")
	}

	return claims.Subject, nil
}
