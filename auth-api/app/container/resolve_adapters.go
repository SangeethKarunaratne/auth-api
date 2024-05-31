package container

import (
	"auth-api/app/config"
	"auth-api/external/adapters"
	"fmt"
)

var resolvedAdapters Adapters

func resolveAdapters(cfg *config.Config) Adapters {

	resolveDBAdapter(cfg.DBConfig)
	resolveLogAdapter(cfg.AppConfig.LoggerConfig)
	return resolvedAdapters
}

func resolveDBAdapter(cfg config.DBConfig) {

	db, err := adapters.NewMySQLAdapter(cfg)
	if err != nil {
		panic(fmt.Sprintf("error: %v", err))
	}

	resolvedAdapters.DBAdapter = db
}

func resolveLogAdapter(cfg config.LoggerConfig) {
	//logger := adapters.NewZapLogger(cfg)
	//resolvedAdapters.LogAdapter = logger
}
