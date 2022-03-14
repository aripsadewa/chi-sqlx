package utils

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

var EnvConfigs *AppConfig

type AppConfig struct {
	AppPort           string        `mapstructure:"APP_PORT"`
	SortCategoryValue string        `mapstructure:"SORT_CATEGORY_VALUE"`
	SecretApp         string        `mapstructure:"SECRET_APP"`
	ExpToken          time.Duration `mapstructure:"EXP_TOKEN"`
	DbUser            string        `mapstructure:"DB_USER"`
	DbPass            string        `mapstructure:"DB_PASS"`
	DbHost            string        `mapstructure:"DB_HOST"`
	DbName            string        `mapstructure:"DB_NAME"`
}

func InitiEnvConfigs() {
	EnvConfigs = loadEnvVariables()
}

func loadEnvVariables() (config *AppConfig) {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error read env file", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}

	return
}
