package tests

import (
	"auth-api/app/container"
	"auth-api/app/http/controllers"
	"auth-api/app/http/response"
	"auth-api/mocks"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

//// Mocking the userUseCase
//// MockUserUseCase is a struct for mocking userUseCase.
//type MockUserUseCase struct {
//	GetUsersFunc func(ctx *gin.Context) ([]entities.User, error)
//}
//
//// GetUsers mocks the GetUsers method.
//func (m *MockUserUseCase) GetUsers(ctx *gin.Context) ([]entities.User, error) {
//	if m.GetUsersFunc != nil {
//		return m.GetUsersFunc(ctx)
//	}
//	return nil, errors.New("GetUsersFunc is not set")
//}

func TestUserController_GetUsersw1(t *testing.T) {
	//ctrl := gomock.NewController(t)
	//mockLogger := mocks.NewMockLoggerInterface(ctrl)
	//mockUserRepo := mocks.NewMockUserRepositoryInterface(ctrl)
	// Set up the container
	//mockContainer := &container.container{
	//	Adapters: container.Adapters{
	//		DBAdapter:  nil,
	//		LogAdapter: mockLogger,
	//	},
	//	Repositories: container.Repositories{
	//		userRepository: mockUserRepo,
	//	},
	//	Services: container.Services{},
	//}
	//mockUserUseCase := mocks.NewMockUserUseCaseInterface(ctrl)
	//
	//// Set up mock behavior for GetUsers
	//mockUserUseCase.GetUsersFunc = func(ctx *gin.Context) ([]entities.User, error) {
	//	return []entities.User{
	//		{ID: 1, Name: "John Doe", Email: "john@example.com"},
	//		{ID: 2, Name: "Jane Doe", Email: "jane@example.com"},
	//	}, nil
	//}
	//
	//userController := mocks.Mockuser_controller_interface{
	//
	//}
	//userController.
	////userController.userUseCase = mockUserUseCase
	//
	//r := gin.Default()
	//r.GET("/users", userController.GetUsers1)
	//
	//req, _ := http.NewRequest(http.MethodGet, "/users", nil)
	//w := httptest.NewRecorder()
	//r.ServeHTTP(w, req)

	//assert.Equal(t, http.StatusOK, w.Code)
	//
	//var usersResponse []response.User
	//err := json.Unmarshal(w.Body.Bytes(), &usersResponse)
	//assert.NoError(t, err)
	//
	//expectedResponse := []response.User{
	//	{ID: 1, Name: "John Doe", Email: "john@example.com"},
	//	{ID: 2, Name: "Jane Doe", Email: "jane@example.com"},
	//}
	//assert.Equal(t, expectedResponse, usersResponse)
}

// Test for GetUsers1 function
func TestUserController_GetUsers1(t *testing.T) {
	// Create a gin engine
	r := gin.Default()

	// Create a new mock use case
	//mockUserUseCase := new(MockUserUseCase)
	ctrl := gomock.NewController(t)
	// Mock data
	//mockUsers := []entities.User{
	//	{ID: 1, Name: "John Doe", Email: "john@example.com"},
	//	{ID: 2, Name: "Jane Doe", Email: "jane@example.com"},
	//}
	//mockUserUseCase.On("GetUsers", mock.Anything).Return(mockUsers, nil)

	mockLogger := mocks.NewMockLoggerInterface(ctrl)
	mockUserRepo := mocks.NewMockUserRepositoryInterface(ctrl)
	// Set up the container
	mockContainer := &container.Container{
		Adapters: container.Adapters{
			DBAdapter:  nil,
			LogAdapter: mockLogger,
		},
		Repositories: container.Repositories{
			UserRepository: mockUserRepo,
		},
		Services: container.Services{},
	}
	userController := controllers.NewUserController(mockContainer)
	//userController.userUseCase = mockUserUseCase

	// Register the GetUsers1 route
	r.GET("/users", userController.GetUsers1)

	// Create a test request
	req, _ := http.NewRequest(http.MethodGet, "/users", nil)
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Validate the response
	assert.Equal(t, http.StatusOK, w.Code)

	var usersResponse []response.User
	err := json.Unmarshal(w.Body.Bytes(), &usersResponse)
	assert.NoError(t, err)

	// Check if the response contains the expected data
	expectedResponse := []response.User{
		{ID: 1, Name: "John Doe", Email: "john@example.com"},
		{ID: 2, Name: "Jane Doe", Email: "jane@example.com"},
	}
	assert.Equal(t, expectedResponse, usersResponse)

	// Ensure that the expectations are met
	//mockUserUseCase.AssertExpectations(t)
}

func SetupTestEnvironment(t *testing.T) (*gomock.Controller, *mocks.MockLoggerInterface, *gin.Engine) {
	// Set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	// Create a new mock controller
	ctrl := gomock.NewController(t)

	// Create a new mock logger
	mockLogger := mocks.NewMockLoggerInterface(ctrl)
	mockUserRepo := mocks.NewMockUserRepositoryInterface(ctrl)

	// Set up expected calls to the mock logger (if any)
	// mockLogger.EXPECT().Info(gomock.Any()).Times(1)

	// Create a new instance of the UserController with the mock container
	mockContainer := &container.Container{
		Adapters: container.Adapters{
			DBAdapter:  nil,
			LogAdapter: mockLogger,
		},
		Repositories: container.Repositories{
			UserRepository: mockUserRepo,
		},
		Services: container.Services{},
	}
	userController := controllers.NewUserController(mockContainer)

	// Create a new Gin router and define the endpoint
	router := gin.Default()
	router.GET("/users", userController.GetUsers1)

	return ctrl, mockLogger, router
}

func TestGetUsers1(t *testing.T) {
	// Set up the test environment
	ctrl, _, router := SetupTestEnvironment(t)
	defer ctrl.Finish()

	// Create a new HTTP request
	req, err := http.NewRequest(http.MethodGet, "/users", nil)
	assert.NoError(t, err)

	// Create a response recorder to record the response
	rr := httptest.NewRecorder()

	// Serve the request
	router.ServeHTTP(rr, req)

	// Check the status code
	assert.Equal(t, http.StatusOK, rr.Code)

	// Define the expected response
	expectedResponse := response.User{
		ID:   1,
		Name: "sddsds",
	}

	// Decode and validate the response
	var actualResponse response.User

	json.NewDecoder(rr.Body).Decode(&actualResponse)

	assert.Equal(t, expectedResponse, actualResponse)
}

func TestGetUsers3(t *testing.T) {
	// Set up the test environment
	ctrl, _, router := SetupTestEnvironment(t)
	defer ctrl.Finish()

	// Create a new HTTP request
	req, err := http.NewRequest(http.MethodGet, "/users", nil)
	assert.NoError(t, err)

	// Create a response recorder to record the response
	rr := httptest.NewRecorder()

	// Serve the request
	router.ServeHTTP(rr, req)

	// Check the status code
	assert.Equal(t, http.StatusOK, rr.Code)

	// Define the expected response
	expectedResponse := response.User{
		ID:   1,
		Name: "sddsds",
	}

	// Decode and validate the response
	var actualResponse response.User

	json.NewDecoder(rr.Body).Decode(&actualResponse)

	assert.Equal(t, expectedResponse, actualResponse)
}

func TestGetUsers2(t *testing.T) {
	// Set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	// Create a new mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a new mock logger
	//mockLogger := mocks.NewMockLoggerInterface(ctrl)

	// Set up expected calls to the mock logger (if any)
	// mockLogger.EXPECT().Info(gomock.Any()).Times(1)

	// Create a new instance of the UserController with the mock container
	//mockContainer := &container.container{
	//	Adapters: container.Adapters{
	//		DBAdapter:  nil,
	//		LogAdapter: mockLogger,
	//	},
	//	Repositories: container.Repositories{},
	//	Services:     container.Services{},
	//}
	//userController := controllers.NewUserController(mockContainer)

	// Create a new Gin router and define the endpoint
	router := gin.Default()
	//router.GET("/users", userController.GetUsers)

	// Create a new HTTP request
	req, err := http.NewRequest(http.MethodGet, "/users", nil)
	assert.NoError(t, err)

	// Create a response recorder to record the response
	rr := httptest.NewRecorder()

	// Serve the request
	router.ServeHTTP(rr, req)

	// Check the status code
	assert.Equal(t, http.StatusOK, rr.Code)

	// Define the expected response
	expectedResponse := response.User{
		ID:   1,
		Name: "sddsds",
	}

	// Decode the response
	var actualResponse response.User
	err = json.NewDecoder(rr.Body).Decode(&actualResponse)
	assert.NoError(t, err)

	// Check if the actual response matches the expected response
	assert.Equal(t, expectedResponse, actualResponse)
}
