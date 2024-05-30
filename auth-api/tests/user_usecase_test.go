package tests

//
//import (
//	"auth-api/domain/entities"
//	"auth-api/domain/usecases"
//	"auth-api/mocks"
//	"context"
//	"go.uber.org/mock/gomock"
//	"testing"
//
//	"github.com/stretchr/testify/assert"
//	"golang.org/x/crypto/bcrypt"
//)
//
//func TestUserUseCase_GetUsers(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	mockRepo := mocks.NewMockUserRepositoryInterface(ctrl)
//	ctx := context.TODO()
//
//	expectedUsers := []entities.User{
//		{ID: 1, Name: "John Doe"},
//	}
//
//	mockRepo.EXPECT().Get(ctx).Return(expectedUsers, nil).Times(1)
//
//	userUseCase := usecases.UserUseCase{userRepository: mockRepo}
//
//	users, err := userUseCase.GetUsers(ctx)
//	assert.NoError(t, err)
//	assert.Equal(t, expectedUsers, users)
//}
//
//func TestUserUseCase_AddUser(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	mockRepo := mocks.NewMockUserRepositoryInterface(ctrl)
//	ctx := context.TODO()
//
//	newUser := entities.User{Email: "john.doe@example.com", Password: "password123"}
//	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
//	expectedUser := entities.User{Email: "john.doe@example.com", Password: string(hashedPassword)}
//
//	mockRepo.EXPECT().UserExists(ctx, newUser.Email).Return(entities.User{}, false, nil).Times(1)
//	mockRepo.EXPECT().Add(ctx, gomock.Any()).DoAndReturn(func(ctx context.Context, user entities.User) error {
//		assert.Equal(t, expectedUser.Email, user.Email)
//		assert.NoError(t, bcrypt.CompareHashAndPassword([]byte(user.Password), []byte("password123")))
//		return nil
//	}).Times(1)
//
//	userUseCase := usecases.UserUseCase{userRepository: mockRepo}
//
//	err := userUseCase.AddUser(ctx, newUser)
//	assert.NoError(t, err)
//}
//
//func TestUserUseCase_AddUser_UserExists(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	mockRepo := mocks.NewMockUserRepositoryInterface(ctrl)
//	ctx := context.TODO()
//
//	existingUser := entities.User{Email: "john.doe@example.com", Password: "password123"}
//
//	mockRepo.EXPECT().UserExists(ctx, existingUser.Email).Return(existingUser, true, nil).Times(1)
//
//	userUseCase := usecases.UserUseCase{userRepository: mockRepo}
//
//	err := userUseCase.AddUser(ctx, existingUser)
//	assert.Error(t, err)
//	assert.Equal(t, "user already registered", err.Error())
//}
//
//func TestUserUseCase_LoginUser(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	mockRepo := mocks.NewMockUserRepositoryInterface(ctrl)
//	ctx := context.TODO()
//
//	userEmail := "john.doe@example.com"
//	userPassword := "password123"
//	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
//	existingUser := entities.User{Email: userEmail, Password: string(hashedPassword)}
//
//	mockRepo.EXPECT().UserExists(ctx, userEmail).Return(existingUser, true, nil).Times(1)
//
//	//userUseCase := usecases.userUseCase{userRepository: mockRepo}
//	//
//	////authenticated, err := userUseCase.LoginUser(ctx, userEmail, userPassword)
//	//assert.NoError(t, err)
//	//assert.True(t, authenticated)
//}
//
//func TestUserUseCase_LoginUser_InvalidPassword(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	//mockRepo := mocks.NewMockUserRepositoryInterface(ctrl)
//	//ctx := context.TODO()
//	//
//	//userEmail := "john.doe@example.com"
//	//userPassword := "password123"
//	//wrongPassword := "wrongpassword"
//	//hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
//	//existingUser := entities.User{Email: userEmail, Password: string(hashedPassword)}
//	//
//	//mockRepo.EXPECT().UserExists(ctx, userEmail).Return(existingUser, true, nil).Times(1)
//	//
//	//userUseCase := usecases.userUseCase{userRepository: mockRepo}
//
//	//authenticated, err := userUseCase.LoginUser(ctx, userEmail, wrongPassword)
//	//assert.Error(t, err)
//	//assert.False(t, authenticated)
//	//assert.Equal(t, "invalid password", err.Error())
//}
//
//func TestUserUseCase_LoginUser_UserNotRegistered(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	//mockRepo := mocks.NewMockUserRepositoryInterface(ctrl)
//	//ctx := context.TODO()
//	//
//	//userEmail := "john.doe@example.com"
//	//userPassword := "password123"
//	//
//	//mockRepo.EXPECT().UserExists(ctx, userEmail).Return(entities.User{}, false, nil).Times(1)
//	//
//	//userUseCase := usecases.userUseCase{userRepository: mockRepo}
//	//
//	//authenticated, err := userUseCase.LoginUser(ctx, userEmail, userPassword)
//	//assert.Error(t, err)
//	//assert.False(t, authenticated)
//	//assert.Equal(t, "user not registered", err.Error())
//}
