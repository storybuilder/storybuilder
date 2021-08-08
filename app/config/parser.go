package config

import (
	"fmt"
	"os"

	yaml "gopkg.in/yaml.v2"
)

// Parse parses all configuration to a single Config object.
// `cfgDir` is the path of the configuration directory.
func Parse(cfgDir string) *Config {
	// set config directory
	dir := getConfigDir(cfgDir)
	return &Config{
		AppConfig:      parseAppConfig(dir),
		DBConfig:       parseDBConfig(dir),
		LogConfig:      parseLogConfig(dir),
		ServicesConfig: parseServicesConfig(dir),
	}
}

// parseAppConfig parses application configurations.
func parseAppConfig(dir string) AppConfig {
	cfg := AppConfig{}
	parseConfig(dir+"app.yaml", &cfg)
	return cfg
}

// parseLogConfig parses logger configurations.
func parseLogConfig(dir string) LogConfig {
	cfg := LogConfig{}
	parseConfig(dir+"logger.yaml", &cfg)
	return cfg
}

// parseDBConfig parses database configurations.
func parseDBConfig(dir string) DBConfig {
	cfg := DBConfig{}
	parseConfig(dir+"database.yaml", &cfg)
	return cfg
}

// parseServicesConfig parses configurations of all services.
func parseServicesConfig(dir string) []ServiceConfig {
	var cfgs []ServiceConfig
	parseConfig(dir+"services.yaml", &cfgs)
	return cfgs
}

// parseConfig reads configuration values from the given file and
// populates the given config struct.
func parseConfig(file string, unpacker interface{}) {
	content := read(file)
	err := yaml.Unmarshal(content, unpacker)
	if err != nil {
		panic(fmt.Sprintf("error: %v", err))
	}
}

// getConfigDir returns config directory path after analysing and correcting.
func getConfigDir(dir string) string {
	// get last char of dir path
	c := dir[len(dir)-1]
	if os.IsPathSeparator(c) {
		return dir
	}
	return dir + string(os.PathSeparator)
}
