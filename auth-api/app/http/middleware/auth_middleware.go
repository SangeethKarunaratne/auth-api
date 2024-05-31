package middleware

import (
	"auth-api/app/container"
	responses "auth-api/app/http/response"
	"auth-api/domain/usecases"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware(ctr *container.Container) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := extractToken(c.GetHeader("Authorization"))

		authUseCase := usecases.NewAuthUseCase(ctr)
		token, err := authUseCase.ValidateToken(c, tokenString)

		if err != nil {
			resp := responses.ErrorResponse{
				Message: "User Unauthorized",
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, resp)
			return
		} else {
			if !token.Valid {
				println(err)
				resp := responses.ErrorResponse{
					Message: "User Unauthorized",
				}
				c.AbortWithStatusJSON(http.StatusUnauthorized, resp)
			}
			return
		}
	}
}

func extractToken(bearerToken string) string {
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}
