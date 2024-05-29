package container

import (
	"auth-api/domain/adapters"
	"auth-api/domain/repositories"
	"auth-api/domain/services"
)

type Container struct {
	Adapters     Adapters
	Repositories Repositories
	Services     Services
}

type Adapters struct {
	DBAdapter  adapters.DBAdapterInterface
	LogAdapter adapters.LoggerInterface
}

type Repositories struct {
	UserRepository repositories.UserRepositoryInterface
}

type Services struct {
	NotificationService services.NotificationServiceAPIInterface
}
