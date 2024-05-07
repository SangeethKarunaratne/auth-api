package container

import "auth-api/app/config"

func Resolve(cfg *config.Config) *Container {

	return &Container{
		Adapters:     resolveAdapters(cfg),
		Services:     resolveServices(cfg),
		Repositories: resolveRepositories(),
	}
}
