package tests

import (
	"bytes"
	"encoding/json"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"

	"auth-api/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestUserControllerInterface_GetUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserController := mocks.NewMockuser_controller_interface(ctrl)

	gin.SetMode(gin.TestMode)
	r := gin.Default()

	r.GET("/users", mockUserController.GetUsers)

	t.Run("successful get users", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/users", nil)
		w := httptest.NewRecorder()

		mockUserController.EXPECT().GetUsers(gomock.Any())

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})
}

func TestUserControllerInterface_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserController := mocks.NewMockuser_controller_interface(ctrl)

	gin.SetMode(gin.TestMode)
	r := gin.Default()

	r.POST("/login", mockUserController.Login)

	t.Run("successful login", func(t *testing.T) {
		loginRequest := map[string]string{
			"email":    "john.doe@example.com",
			"password": "password123",
		}
		reqBody, _ := json.Marshal(loginRequest)
		req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		mockUserController.EXPECT().Login(gomock.Any())

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusAccepted, w.Code)
	})

	t.Run("login invalid request", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer([]byte("invalid request")))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		mockUserController.EXPECT().Login(gomock.Any())

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestUserControllerInterface_Register(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserController := mocks.NewMockuser_controller_interface(ctrl)

	gin.SetMode(gin.TestMode)
	r := gin.Default()

	r.POST("/register", mockUserController.Register)

	t.Run("successful registration", func(t *testing.T) {
		registerRequest := map[string]string{
			"name":     "John Doe",
			"email":    "john.doe@example.com",
			"password": "password123",
		}
		reqBody, _ := json.Marshal(registerRequest)
		req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		mockUserController.EXPECT().Register(gomock.Any())

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusAccepted, w.Code)
	})

	t.Run("register invalid request", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer([]byte("invalid request")))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		mockUserController.EXPECT().Register(gomock.Any())

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
