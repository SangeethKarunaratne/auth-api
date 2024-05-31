package usecases

import (
	"auth-api/app/container"
	error2 "auth-api/app/http/error"
	"auth-api/domain/entities"
	"auth-api/domain/repositories"
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

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
		return nil, err
	}

	return users, nil
}

func (s *UserUseCase) AddUser(ctx context.Context, user entities.User) error {

	_, isUserExists, err := s.userRepository.UserExists(ctx, user.Email)

	if err != nil {
		return err
	}

	if isUserExists {
		return errors.New("user already registered")
	} else {
		user.Password, err = s.generateHashPassword(ctx, user.Password)
		if err != nil {
			return errors.New("an error occurred")
		}

		err = s.userRepository.Add(ctx, user)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *UserUseCase) generateHashPassword(ctx context.Context, pass string) (passHash string, err error) {

	passHashBytes, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	passHash = string(passHashBytes)
	return passHash, nil
}

func (s *UserUseCase) LoginUser(ctx context.Context, email string, password string) (authenticated bool, err error) {

	user, isUserExists, err := s.userRepository.UserExists(ctx, email)

	if err != nil {
		return false, err
	}

	if !isUserExists {
		return false, error2.ErrInvalidCredentials
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		// todo: add log here
		return false, error2.ErrInvalidCredentials
	}

	return true, nil
}
