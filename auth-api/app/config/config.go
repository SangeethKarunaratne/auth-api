package config

type Config struct {
	AppConfig AppConfig
	DBConfig  DBConfig
}

type AppConfig struct {
	Name                      string                    `yaml:"name"`
	Port                      int                       `yaml:"port"`
	NotificationServiceConfig NotificationServiceConfig `yaml:"notification_service_config"`
	LoggerConfig              LoggerConfig              `yaml:"logger_config"`
}

type DBConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Database string `yaml:"database"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}
type NotificationServiceConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type LoggerConfig struct {
	Level string `yaml:"level"`
}
