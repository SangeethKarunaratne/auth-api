package repositories

import (
	"auth-api/domain/entities"
	"context"
)

type UserRepositoryInterface interface {
	Get(ctx context.Context) ([]entities.User, error)
	GetByID(ctx context.Context, id int) (entities.User, error)
	UserExists(ctx context.Context, email string) (entities.User, bool, error)
	Add(ctx context.Context, sample entities.User) error
	Edit(ctx context.Context, sample entities.User) error
	Delete(ctx context.Context, id int) error
}
