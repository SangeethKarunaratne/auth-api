package usecases

import (
	"auth-api/app/container"
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

const authenticationUseCaseLogPrefix = "AuthenticationUseCase"

type AuthUseCase struct {
	secretKey string
	issuer    string
}

type authCustomClaims struct {
	Name string `json:"name"`
	User bool   `json:"user"`
	jwt.StandardClaims
}

func NewAuthUseCase(container *container.Container) *AuthUseCase {
	return &AuthUseCase{
		secretKey: getSecretKey(),
		issuer:    "AUTH-API",
	}
}

func (service *AuthUseCase) GenerateToken(ctx context.Context, email string, isUser bool) string {
	claims := &authCustomClaims{
		email,
		isUser,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    service.issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		panic(err)
	}

	return t
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
