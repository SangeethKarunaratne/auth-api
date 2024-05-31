package usecases

import (
	"auth-api/app/container"
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

type AuthUseCase struct {
	secretKey string
	issuer    string
}

type authCustomClaims struct {
	Name string `json:"name"`
	User bool   `json:"user"`
	jwt.RegisteredClaims
}

func NewAuthUseCase(container *container.Container) *AuthUseCase {
	return &AuthUseCase{
		secretKey: getSecretKey(),
		issuer:    "AUTH-API",
	}
}

func (service *AuthUseCase) GenerateToken(ctx context.Context, email string, isUser bool) (string, error) {
	claims := &authCustomClaims{
		email,
		isUser,
		jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			Issuer:    service.issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		return "", err
	}

	return t, err
}

func (service *AuthUseCase) ValidateToken(ctx context.Context, encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("Invalid Token", token.Header["alg"])
		}
		return []byte(service.secretKey), nil
	})
}

func getSecretKey() string {
	secret := os.Getenv("SECRET")
	if secret == "" {
		secret = "secret"
	}
	return secret
}
