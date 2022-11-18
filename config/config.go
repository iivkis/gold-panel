package config

import "github.com/spf13/viper"

type Config struct {
	Panel struct {
		BotToken string `mapstructure:"bot_token"`
	} `mapstructure:"panel"`
}

var config Config

func LoadFrom(path string) {
	viper.AddConfigPath(path)

	viper.SetConfigType("yaml")
	viper.SetConfigName("config.yaml")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		panic(err)
	}
}

func Get() *Config {
	return &config
}
