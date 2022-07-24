package container

import (
	"auth-api/domain/adapters"
	"auth-api/domain/repositories"
)

type Container struct {
	Adapters     Adapters
	Repositories Repositories
}

type Adapters struct {
	DBAdapter adapters.DBAdapterInterface
}

type Repositories struct {
	UserRepository repositories.UserRepositoryInterface
}
