package repositories

import (
	"auth-api/domain/adapters"
	"auth-api/domain/entities"
	"auth-api/domain/repositories"
	"context"
)

type UserRepository struct {
	db adapters.DBAdapterInterface
}

func NewUserRepository(dbAdapter adapters.DBAdapterInterface) repositories.UserRepositoryInterface {
	return &UserRepository{db: dbAdapter}
}

func (u UserRepository) Get(ctx context.Context) ([]entities.User, error) {
	query := `SELECT id, name, email
				FROM users`

	parameters := map[string]interface{}{}

	result, err := u.db.Query(ctx, query, parameters)
	if err != nil {
		return nil, err
	}

	return u.mapResult(result), nil
}

func (u UserRepository) UserExists(ctx context.Context, email string) (entities.User, bool, error) {
	query := `SELECT id, name, password
				FROM users
				WHERE email=?email limit 1`

	parameters := map[string]interface{}{
		"email": email,
	}

	result, err := u.db.Query(ctx, query, parameters)
	if err != nil {
		return entities.User{}, false, err
	}

	mapped := u.mapResult(result)
	if len(mapped) == 0 {
		return entities.User{}, false, nil
	}

	return mapped[0], true, nil
}

func (u UserRepository) Add(ctx context.Context, user entities.User) error {
	query := `INSERT INTO users
				(name,email, password)
				VALUES(?name,?email,?password)
				`

	parameters := map[string]interface{}{
		"name":     user.Name,
		"email":    user.Email,
		"password": user.Password,
	}

	_, err := u.db.Query(ctx, query, parameters)
	if err != nil {
		return err
	}

	return nil
}

func (u UserRepository) GetByID(ctx context.Context, id int) (entities.User, error) {
	panic("implement me")
}

func (u UserRepository) Edit(ctx context.Context, user entities.User) error {
	panic("implement me")
}

func (u UserRepository) Delete(ctx context.Context, id int) error {
	panic("implement me")
}
func (u UserRepository) mapResult(result []map[string]interface{}) []entities.User {

	m := make([]entities.User, 0)

	for _, row := range result {

		id, _ := row["id"].(int64)
		name, _ := row["name"].([]byte)
		email, _ := row["email"].([]byte)
		password, _ := row["password"].([]byte)

		m = append(m, entities.User{
			ID:       int(id),
			Name:     string(name),
			Email:    string(email),
			Password: string(password),
		})
	}

	return m
}
