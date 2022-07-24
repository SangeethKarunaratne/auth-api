package usecases

import (
	"auth-api/app/container"
	"auth-api/domain/entities"
	"auth-api/domain/repositories"
	"context"
	"errors"
	"fmt"
	"github.com/pickme-go/log"
	"golang.org/x/crypto/bcrypt"
)

const userUseCaseLogPrefix = "UserUseCase"

type UserUseCase struct {
	userRepository repositories.UserRepositoryInterface
}

func NewUserUseCase(container *container.Container) *UserUseCase {
	return &UserUseCase{
		userRepository: container.Repositories.UserRepository,
	}
}

func (s *UserUseCase) GetUsers(ctx context.Context) ([]entities.User, error) {

	users, err := s.userRepository.Get(ctx)
	if err != nil {
		log.ErrorContext(ctx, userUseCaseLogPrefix, err)
		return nil, err
	}

	return users, nil
}

func (s *UserUseCase) AddUser(ctx context.Context, user entities.User) error {

	_, isUserExists, err := s.userRepository.UserExists(ctx, user.Email)

	if err != nil {
		log.ErrorContext(ctx, userUseCaseLogPrefix, err)
		return err
	}

	if isUserExists {
		log.InfoContext(ctx, userUseCaseLogPrefix, fmt.Sprintf("user already registered with email: %v", user.Email))
		return errors.New("user already registered")
	} else {
		user.Password, err = s.generateHashPassword(ctx, user.Password)
		if err != nil {
			log.ErrorContext(ctx, userUseCaseLogPrefix, fmt.Sprintf("password hashing failed err: %v", err))
			return errors.New("an error occurred")
		}

		err = s.userRepository.Add(ctx, user)
		if err != nil {
			log.ErrorContext(ctx, userUseCaseLogPrefix, err)
			return err
		}
	}

	return nil
}

func (s *UserUseCase) generateHashPassword(ctx context.Context, pass string) (passHash string, err error) {

	passHashBytes, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		log.ErrorContext(ctx, userUseCaseLogPrefix, fmt.Sprintf("password hashing failed err: %v", err))
		return "", err
	}
	passHash = string(passHashBytes)
	return passHash, nil
}

func (s *UserUseCase) LoginUser(ctx context.Context, email string, password string) (authenticated bool, err error) {

	user, isUserExists, err := s.userRepository.UserExists(ctx, email)

	if err != nil {
		log.ErrorContext(ctx, userUseCaseLogPrefix, fmt.Sprintf("err: %v", err))
		return false, err
	}

	if !isUserExists {
		log.InfoContext(ctx, userUseCaseLogPrefix, "user not registered")
		return false, errors.New("user not registered")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		log.InfoContext(ctx, userUseCaseLogPrefix, err)
		return false, errors.New("invalid password")
	}

	return true, nil
}
