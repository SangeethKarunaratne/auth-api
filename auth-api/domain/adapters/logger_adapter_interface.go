package adapters

type LoggerInterface interface {
	Warn(message interface{}, params ...interface{})
	Info(message interface{}, params ...interface{})
	Error(message interface{}, params ...interface{})
	Debug(message interface{}, params ...interface{})
}
