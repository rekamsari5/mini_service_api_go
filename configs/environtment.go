package configs

import "github.com/spf13/viper"

var Configs *Config

type Config struct {
	APP_NAME      string `mapstructure:"APP_NAME"`
	PORT          string `mapstructure:"PORT"`
	GIN_MODE      string `mapstructure:"GIN_MODE"`
	APP_ENV       string `mapstructure:"APP_ENV"`
	TIMEOUT_REQ   int64  `mapstructure:"TIMEOUT_REQUEST"`
	LOCAL         string `mapstructure:"LOCAL"`
	DEBUG         bool   `mapstructure:"DEBUG"`
	DB_USERNAME   string `mapstructure:"DB_USERNAME"`
	DB_PASSWORD   string `mapstructure:"DB_PASSWORD"`
	DB_HOST       string `mapstructure:"DB_HOST"`
	DB_PORT       string `mapstructure:"DB_PORT"`
	DB_DATABASE   string `mapstructure:"DB_DATABASE"`
	TIME_LOCATION string `mapstructure:"TIME_LOCATION"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("dev")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	Configs = &config
	return
}
