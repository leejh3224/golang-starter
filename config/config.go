package config

import (
	"os"

	"github.com/spf13/viper"
)

// LoadConfig ...
func LoadConfig() *viper.Viper {
	v := viper.New()
	v.SetConfigType("json")
	v.SetConfigName("config." + getEnv("APP_ENV", "development"))
	v.AddConfigPath(".")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	return v
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
