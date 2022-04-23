package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func CheckIfSet(key string) {
	if !viper.IsSet(key) {
		err := fmt.Errorf("key %s is not set", key)
		panic(err)
	}
}

func ReadEnvString(key string) string {
	CheckIfSet(key)
	return viper.GetString(key)
}

func ReadEnvInt(key string) int {
	CheckIfSet(key)
	return viper.GetInt(key)
}
