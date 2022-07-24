package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

func Parse(cfgDir string) *Config {
	dir := getConfigDir(cfgDir)
	return &Config{
		AppConfig: parseAppConfig(dir),
		DBConfig:  parseDBConfig(dir),
	}
}

func parseAppConfig(dir string) AppConfig {
	cfg := AppConfig{}
	parseConfig(dir+"app.yaml", &cfg)
	return cfg
}

func parseDBConfig(dir string) DBConfig {
	cfg := DBConfig{}
	parseConfig(dir+"database.yaml", &cfg)
	return cfg
}

func parseConfig(file string, unpacker interface{}) {

	content := read(file)

	err := yaml.Unmarshal(content, unpacker)
	if err != nil {
		panic(fmt.Sprintf("error: %v", err))
	}
}

func getConfigDir(dir string) string {

	c := dir[len(dir)-1]
	if os.IsPathSeparator(c) {

		return dir
	}

	return dir + string(os.PathSeparator)
}
