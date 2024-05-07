package container

import (
	"auth-api/app/config"
	"auth-api/external/services"
)

var resolvedServices Services

func resolveServices(cfg *config.Config) Services {

	resolveNotificationService(cfg.AppConfig.NotificationServiceConfig)
	return resolvedServices
}

func resolveNotificationService(cfg config.NotificationServiceConfig) {

	notificationService := services.NewNotificationServiceAPI(cfg)
	resolvedServices.NotificationService = notificationService
}
