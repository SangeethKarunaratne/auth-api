package tests

import (
	"auth-api/app/config"
	"auth-api/app/container"
	"auth-api/app/http/controllers"
	"auth-api/app/http/request"
	"auth-api/app/http/response"
	"auth-api/domain/entities"
	"auth-api/external/adapters"
	"auth-api/external/repositories"
	"auth-api/mocks"
	"bytes"
	"encoding/json"
	"errors"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUserRepo := mocks.NewMockUserRepositoryInterface(mockCtrl)
	mockLogger := mocks.NewMockLoggerInterface(mockCtrl)

	container := &container.Container{
		Repositories: container.Repositories{
			UserRepository: mockUserRepo,
		},
		Adapters: container.Adapters{
			LogAdapter: mockLogger,
		},
	}

	userController := controllers.NewUserController(container)

	t.Run("success", func(t *testing.T) {
		reqBody := request.RegisterRequest{
			Name:     "John Doe",
			Email:    "john@example.com",
			Password: "password",
		}

		reqBodyBytes, _ := json.Marshal(reqBody)
		req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(reqBodyBytes))
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(rr)
		ctx.Request = req

		user := entities.User{
			Name:     reqBody.Name,
			Email:    reqBody.Email,
			Password: reqBody.Password,
		}

		mockUserRepo.EXPECT().UserExists(gomock.Any(), user.Email).Return(entities.User{}, false, nil)
		mockUserRepo.EXPECT().Add(gomock.Any(), gomock.Any()).Return(nil)

		userController.Register(ctx)

		assert.Equal(t, http.StatusAccepted, rr.Code)
		assert.JSONEq(t, `{"message": "user successfully created"}`, rr.Body.String())
	})

	t.Run("validation failure", func(t *testing.T) {
		reqBody := `{"Name": "", "Email": "invalid", "Password": "short"}`

		req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer([]byte(reqBody)))
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(rr)
		ctx.Request = req

		userController.Register(ctx)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("use case failure", func(t *testing.T) {
		reqBody := request.RegisterRequest{
			Name:     "John Doe",
			Email:    "john@example.com",
			Password: "password",
		}

		reqBodyBytes, _ := json.Marshal(reqBody)
		req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(reqBodyBytes))
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(rr)
		ctx.Request = req

		user := entities.User{
			Name:     reqBody.Name,
			Email:    reqBody.Email,
			Password: reqBody.Password,
		}

		mockUserRepo.EXPECT().UserExists(gomock.Any(), user.Email).Return(entities.User{}, false, nil)
		mockUserRepo.EXPECT().Add(gomock.Any(), gomock.Any()).Return(errors.New("use case error"))

		userController.Register(ctx)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.JSONEq(t, `{"message": "use case error"}`, rr.Body.String())
	})
}

func TestGetUsers(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUserRepo := mocks.NewMockUserRepositoryInterface(mockCtrl)
	mockLogger := mocks.NewMockLoggerInterface(mockCtrl)
	mockDB := mocks.NewMockDBAdapterInterface(mockCtrl)

	container := &container.Container{
		Repositories: container.Repositories{
			UserRepository: mockUserRepo,
		},
		Adapters: container.Adapters{
			LogAdapter: mockLogger,
		},
	}

	userController := controllers.NewUserController(container)
	userRep := repositories.NewUserRepository(mockDB)

	t.Run("success", func(t *testing.T) {
		users := []entities.User{
			{ID: 1, Name: "John Doe", Email: "john@example.com"},
			{ID: 2, Name: "Jane Doe", Email: "jane@example.com"},
		}

		var queryResponse []map[string]interface{}
		newMap := map[string]interface{}{
			"name":  "John Doe",
			"age":   30,
			"email": "john.doe@example.com",
		}
		queryResponse = append(queryResponse, newMap)

		rr := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(rr)

		mockDB.EXPECT().Query(ctx, `SELECT id, name, email FROM users`,
			map[string]interface{}{}).Return(queryResponse, nil)
		mockUserRepo.EXPECT().Get(gomock.Any()).Return(users, nil)
		userRep.Get(ctx)

		userController.GetUsers(ctx)

		assert.Equal(t, http.StatusOK, rr.Code)

		expectedResponse, _ := json.Marshal([]response.User{
			{ID: 1, Name: "John Doe", Email: "john@example.com"},
			{ID: 2, Name: "Jane Doe", Email: "jane@example.com"},
		})
		assert.JSONEq(t, string(expectedResponse), rr.Body.String())
	})

	t.Run("use case failure", func(t *testing.T) {
		rr := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(rr)

		mockUserRepo.EXPECT().Get(gomock.Any()).Return(nil, errors.New("use case error"))

		userController.GetUsers(ctx)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.JSONEq(t, `{"message": "use case error"}`, rr.Body.String())
	})
}

func TestLogin(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUserRepo := mocks.NewMockUserRepositoryInterface(mockCtrl)
	mockLogger := mocks.NewMockLoggerInterface(mockCtrl)

	container := &container.Container{
		Repositories: container.Repositories{
			UserRepository: mockUserRepo,
		},
		Adapters: container.Adapters{
			LogAdapter: mockLogger,
		},
	}
	logger := adapters.NewZapLogger(config.LoggerConfig{Level: "info"})

	// Create UserController instance
	userController := controllers.NewUserController(container)

	t.Run("success", func(t *testing.T) {

		loginRequest := request.LoginRequest{
			Email:    "john@example.com",
			Password: "mysecretpassword",
		}

		reqBodyBytes, _ := json.Marshal(loginRequest)
		req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(reqBodyBytes))
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(rr)
		ctx.Set("logger", logger)
		ctx.Request = req

		mockUserRepo.EXPECT().UserExists(gomock.Any(), loginRequest.Email).Return(entities.User{
			ID:       0,
			Name:     "john",
			Password: "$2b$12$7KrfnbnfIsEy3bK6Y.f/4uICIcohkln2dG5aBVzf9dBruWF9J6QmC",
			Email:    "john@example.com",
		}, true, nil)

		userController.Login(ctx)

		assert.Equal(t, http.StatusAccepted, rr.Code)
		var tokenResp response.TokenResponse
		json.Unmarshal(rr.Body.Bytes(), &tokenResp)
		assert.NotEmpty(t, tokenResp.Token)
	})

	t.Run("validation failure", func(t *testing.T) {
		reqBody := `{"Email": "invalid", "Password": ""}`

		req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer([]byte(reqBody)))
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(rr)
		ctx.Set("logger", logger)
		ctx.Request = req

		userController.Login(ctx)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("authentication failure", func(t *testing.T) {

		loginRequest := request.LoginRequest{
			Email:    "john@example.com",
			Password: "wrongpassword",
		}

		reqBodyBytes, _ := json.Marshal(loginRequest)
		req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(reqBodyBytes))
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(rr)
		ctx.Set("logger", logger)
		ctx.Request = req

		mockUserRepo.EXPECT().UserExists(gomock.Any(), loginRequest.Email).Return(entities.User{
			ID:       0,
			Name:     "john",
			Password: "$2b$12$7KrfnbnfIsEy3bK6Y.f/4uICIcohkln2dG5aBVzf9dBruWF9J6QmC",
			Email:    "john@example.com",
		}, true, nil)

		userController.Login(ctx)

		assert.Equal(t, http.StatusUnauthorized, rr.Code)
		assert.JSONEq(t, `{"message": "invalid credentials"}`, rr.Body.String())
	})
}
