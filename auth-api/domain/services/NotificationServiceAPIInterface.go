package services

type NotificationServiceAPIInterface interface {
	SendEmailNotification(email string, content []byte) error
	SendSMSNotification(mobileNumber string, content []byte) error
	SendSlackNotification(channel string, content []byte) error
}
