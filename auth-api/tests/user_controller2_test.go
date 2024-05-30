package tests

import (
	"auth-api/app/http/response"
	"auth-api/mocks"
	"bytes"
	"github.com/gin-gonic/gin"
	"go.uber.org/mock/gomock"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestLogin_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	//mockUserUseCase := mocks.NewMockUserRepositoryInterface(ctrl)
	//mockUserUseCase := mocks.NewMockUserUseCase(ctrl)
	//mockUserRepository := mocks.NewMockUserRepositoryInterface(ctrl)
	//
	//ctx := context.Background()
	//ctx := context.Background()
	userController := mocks.NewMockuser_controller_interface(ctrl)
	var userResponse response.User
	userResponse.ID = 1
	userResponse.Name = "sddsds"
	//mockContext, _ := gin.CreateTestContext(nil)
	mockContext, _ := gin.CreateTestContext(nil)
	requestBody := []byte(`{"email":"test@example.com","password":"password"}`)
	mockContext.Request = &http.Request{}
	mockContext.Request.Body = ioutil.NopCloser(bytes.NewReader(requestBody))
	userController.EXPECT().GetUsers(mockContext).Return(userResponse, true, nil)

	// Mock the input request and expected user credentials
	//email := "test@example.com"
	////password := "password"
	//
	//requestBody := []byte(`{"email":"test@example.com","password":"password"}`)
	//mockContext.Request = &http.Request{}
	//mockContext.Request.Body = ioutil.NopCloser(bytes.NewReader(requestBody))
	//
	////mockContext, _ := gin.CreateTestContext(nil)
	////mockContext.Request = &http.Request{}
	////mockContext.Request.Body = gin.HTTPBody{([]byte)(`{"email":"test@example.com","password":"password"}`)}
	//
	//// Mock the user exists check
	//mockUserRepository.EXPECT().UserExists(ctx, email).Return(entities.User{}, true, nil)
	//
	//// Mock the password comparison
	////mockUserUseCase.EXPECT().(ctx, email, password).Return(true, nil)
	//
	//// Call the login method
	////userController.Login(mockContext)
	//
	//// Assert the status code
	//if mockContext.Writer.Status() != http.StatusAccepted {
	//	t.Errorf("Expected status code %d, got %d", http.StatusAccepted, mockContext.Writer.Status())
	//}
}

//
//func TestLogin_InvalidCredentials(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	mockUserUseCase := mocks.NewMockUserUseCase(ctrl)
//	mockUserRepository := mocks.NewMockUserRepositoryInterface(ctrl)
//
//	ctx := context.Background()
//	userController := controllers.NewUserController(mockUserUseCase, mockUserRepository)
//
//	// Mock the input request with invalid credentials
//	mockContext, _ := gin.CreateTestContext(nil)
//	mockContext.Request = &http.Request{}
//	mockContext.Request.Body = gin.HTTPBody{([]byte)(`{"email":"test@example.com","password":"wrongpassword"}`)}
//
//	// Mock the user exists check
//	mockUserRepository.EXPECT().UserExists(ctx, "test@example.com").Return(entities.User{}, false, nil)
//
//	// Call the login method
//	userController.Login(mockContext)
//
//	// Assert the status code
//	if mockContext.Writer.Status() != http.StatusUnauthorized {
//		t.Errorf("Expected status code %d, got %d", http.StatusUnauthorized, mockContext.Writer.Status())
//	}
//}
//
//func TestLogin_UserNotRegistered(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	mockUserUseCase := mocks.NewMockUserUseCase(ctrl)
//	mockUserRepository := mocks.NewMockUserRepositoryInterface(ctrl)
//
//	ctx := context.Background()
//	userController := controllers.NewUserController(mockUserUseCase, mockUserRepository)
//
//	// Mock the input request with unregistered user
//	mockContext, _ := gin.CreateTestContext(nil)
//	mockContext.Request = &http.Request{}
//	mockContext.Request.Body = gin.HTTPBody{([]byte)(`{"email":"test@example.com","password":"password"}`)}
//
//	// Mock the user exists check
//	mockUserRepository.EXPECT().UserExists(ctx, "test@example.com").Return(entities.User{}, false, nil)
//
//	// Call the login method
//	userController.Login(mockContext)
//
//	// Assert the status code
//	if mockContext.Writer.Status() != http.StatusUnauthorized {
//		t.Errorf("Expected status code %d, got %d", http.StatusUnauthorized, mockContext.Writer.Status())
//	}
//}
//
//func TestRegister_Success(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	mockUserUseCase := mocks.NewMockUserUseCase(ctrl)
//	mockUserRepository := mocks.NewMockUserRepositoryInterface(ctrl)
//
//	ctx := context.Background()
//	userController := controllers.NewUserController(mockUserUseCase, mockUserRepository)
//
//	// Mock the input request and expected user data
//	mockContext, _ := gin.CreateTestContext(nil)
//	mockContext.Request = &http.Request{}
//	mockContext.Request.Body = gin.HTTPBody{([]byte)(`{"name":"Test User","email":"test@example.com","password":"password"}`)}
//	newUser := entities.User{Name: "Test User", Email: "test@example.com", Password: "password"}
//
//	// Mock the user existence check
//	mockUserRepository.EXPECT().UserExists(ctx, "test@example.com").Return(entities.User{}, false, nil)
//
//	// Mock the user addition
//	mockUserUseCase.EXPECT().AddUser(ctx, newUser).Return(nil)
//
//	// Call the register method
//	userController.Register(mockContext)
//
//	// Assert the status code
//	if mockContext.Writer.Status() != http.StatusAccepted {
//		t.Errorf("Expected status code %d, got %d", http.StatusAccepted, mockContext.Writer.Status())
//	}
//}
//
//func TestGetUsers_Success(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	mockUserUseCase := mocks.NewMockUserUseCase(ctrl)
//	mockUserRepository := mocks.NewMockUserRepositoryInterface(ctrl)
//
//	userController := mocks.NewMockuser_controller_interface(ctrl)
//
//	ctx := context.Background()
//	//userController := controllers.NewUserController(mockUserUseCase, mockUserRepository)
//
//	// Mock the user list
//	users := []entities.User{
//		{ID: 1, Name: "User 1", Email: "user1@example.com"},
//		{ID: 2, Name: "User 2", Email: "user2@example.com"},
//	}
//
//	// Mock the user retrieval
//	mockUserUseCase.EXPECT().GetUsers(ctx).Return(users, nil)
//
//	// Mock the response writer
//	mockContext, _ := gin.CreateTestContext(nil)
//	mockContext.Request = &http.Request{}
//
//	// Call the get users method
//	userController.GetUsers(mockContext)
//
//	// Assert the status code
//	if mockContext.Writer.Status() != http.StatusOK {
//		t.Errorf("Expected status code %d, got %d", http.StatusOK, mockContext.Writer.Status())
//	}
//
//	// Assert the response body
//	// You can further assert the response body if needed
//}
