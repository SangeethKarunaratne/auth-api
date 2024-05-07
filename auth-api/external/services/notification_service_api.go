package services

import (
	"auth-api/app/config"
	"auth-api/domain/services"
)

type NotificationServiceAPI struct {
	cfg config.NotificationServiceConfig
}

func NewNotificationServiceAPI(cfg config.NotificationServiceConfig) services.NotificationServiceAPIInterface {
	a := &NotificationServiceAPI{
		cfg: cfg,
	}
	return a
}

func (ns *NotificationServiceAPI) SendEmailNotification(email string, content []byte) error {
	return nil
}

func (ns *NotificationServiceAPI) SendSMSNotification(mobileNumber string, content []byte) error {
	return nil
}

func (ns *NotificationServiceAPI) SendSlackNotification(channel string, content []byte) error {
	return nil
}
