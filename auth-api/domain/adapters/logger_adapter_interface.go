package adapters

type LoggerInterface interface {
	Warn1(message interface{}, params ...interface{})
	Info1(message interface{}, params ...interface{})
	Error1(message interface{}, params ...interface{})
	Debug1(message interface{}, params ...interface{})
}
