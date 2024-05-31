package error

import (
	"auth-api/app/http/response"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type ValidationErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func GetErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "lte":
		return "Should be less than " + fe.Param()
	case "gte":
		return "Should be greater than " + fe.Param()
	}
	return "Unknown error"
}

var ErrInvalidCredentials = errors.New("invalid credentials")

func HandleClientError(ctx *gin.Context, statusCode int, err any) {
	ctx.AbortWithStatusJSON(statusCode, err)
}

func HandleServerError(ctx *gin.Context, err error) {
	// Log the internal server error for internal purposes
	fmt.Printf("Internal server error: %v\n", err)

	// Mask the internal error with a generic message
	ctx.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrorResponse{
		Message: "Internal server error",
	})
}
