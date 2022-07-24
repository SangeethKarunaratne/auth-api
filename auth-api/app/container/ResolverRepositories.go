package container

import "auth-api/external/repositories"

var resolvedRepositories Repositories

func resolveRepositories() Repositories {

	resolvedRepositories.UserRepository = repositories.NewUserRepository(resolvedAdapters.DBAdapter)
	return resolvedRepositories
}
