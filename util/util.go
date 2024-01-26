package util

import (
	"github.com/spf13/viper"
)

type Envs struct {
	ApiKey string `mapstructure:"API_KEY"`
}

func LoadEnvs(path string) (config *Envs, errs error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	viper.AutomaticEnv()
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}
	return
}