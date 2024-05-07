package controllers

import (
	"auth-api/app/container"
	_ "auth-api/app/http/error"
	error2 "auth-api/app/http/error"
	"auth-api/app/http/request"
	"auth-api/app/http/response"
	_ "auth-api/app/http/response"
	"auth-api/domain/entities"
	"auth-api/domain/usecases"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pickme-go/log"
	"net/http"
)

const logPrefix = "UserController"

type UserController struct {
	container   *container.Container
	userUseCase *usecases.UserUseCase
	authUseCase *usecases.AuthUseCase
}

func NewUserController(container *container.Container) *UserController {
	return &UserController{
		container:   container,
		userUseCase: usecases.NewUserUseCase(container),
		authUseCase: usecases.NewAuthUseCase(container),
	}
}

func (controller *UserController) Login(ctx *gin.Context) {

	loginRequest := request.LoginRequest{}
	if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]error2.ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = error2.ErrorMsg{Field: fe.Field(), Message: error2.GetErrorMsg(fe)}
			}
			log.InfoContext(ctx, logPrefix, out)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
			return
		}
	}

	isUserAuthenticated, err := controller.userUseCase.LoginUser(ctx, loginRequest.Email, loginRequest.Password)

	if err != nil {
		log.ErrorContext(ctx, logPrefix, err)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	if isUserAuthenticated {
		token := controller.authUseCase.GenerateToken(ctx, loginRequest.Email, true)
		ctx.IndentedJSON(http.StatusAccepted, response.TokenResponse{
			Token: token,
		})
		log.InfoContext(ctx, logPrefix, fmt.Sprintf("Token generate for user: %v, token: %v", loginRequest.Email, token))
		return
	} else {
		log.InfoContext(ctx, logPrefix, "incorrect email or password")
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "incorrect email or password"})
		return
	}

}

func (controller *UserController) Register(ctx *gin.Context) {

	RegisterRequest := request.RegisterRequest{}
	if err := ctx.ShouldBindJSON(&RegisterRequest); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]error2.ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = error2.ErrorMsg{Field: fe.Field(), Message: error2.GetErrorMsg(fe)}
			}
			log.InfoContext(ctx, out)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
			return
		}
	}

	user := entities.User{
		Name:     RegisterRequest.Name,
		Password: RegisterRequest.Password,
		Email:    RegisterRequest.Email,
	}

	err := controller.userUseCase.AddUser(ctx, user)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusAccepted, gin.H{"message": "user successfully created"})
	return
}

func (controller *UserController) GetUsers(ctx *gin.Context) {

	users, err := controller.userUseCase.GetUsers(ctx)

	var userResponse []response.User

	for _, obj := range users {
		ur := response.User{
			ID:    obj.ID,
			Name:  obj.Name,
			Email: obj.Email,
		}
		userResponse = append(userResponse, ur)
	}

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, userResponse)

}
