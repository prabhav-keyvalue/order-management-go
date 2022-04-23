package config

import (
	"github.com/spf13/viper"
)

type Environment string

const (
	Development Environment = "development"
	Production  Environment = "production"
	Staging     Environment = "staging"
)

type config struct {
	Env  Environment
	Port string

	// db config
	databaseConfig databaseConfig
}

var appConfig config

func Load(filePath string) (err error) {
	viper.SetDefault("ENV", Development)
	viper.AutomaticEnv()
	viper.SetConfigFile(filePath)

	err = viper.ReadInConfig()

	if _, ok := err.(viper.ConfigFileNotFoundError); err != nil && !ok {
		return
	}

	appConfig = config{
		Port:           ReadEnvString("PORT"),
		Env:            Environment(ReadEnvString("ENV")),
		databaseConfig: newDatabaseConfig(),
	}
	return
}

func GetAppPort() string {
	return appConfig.Port
}

func GetEnv() string {
	return string(appConfig.Env)
}

func GetDbConfig() databaseConfig {
	return appConfig.databaseConfig
}
