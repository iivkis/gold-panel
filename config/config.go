package config

import "github.com/spf13/viper"

type Config struct {
	Panel struct {
		BotToken string `mapstructure:"bot_token"`
	} `mapstructure:"panel"`
}

type Env struct {
	MySQLUser     string `mapstructure:"MYSQL_USER"`
	MySQLPassword string `mapstructure:"MYSQL_PASSWORD"`
	MySQLDBName   string `mapstructure:"MYSQL_DBNAME"`
}

var (
	config *Config
	env    *Env
)

func LoadConfigFrom(path string) {
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

func LoadEnvFrom(path string) {
	viper.AddConfigPath(path)

	viper.SetConfigType("env")
	viper.SetConfigName(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&env); err != nil {
		panic(err)
	}
}

func GetConfig() *Config {
	return config
}

func GetEnv() *Env {
	return env
}
