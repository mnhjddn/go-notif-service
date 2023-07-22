package infra

import "github.com/spf13/viper"

// config struct

type AppConfig struct {
	App      App
	RabbitMQ RabbitMQ
	MySQL    MySQL
}

type App struct {
	Env  string
	Name string
	Host string
	Port string
}

type RabbitMQ struct {
	Host       string
	Port       string
	Usr        string
	Pw         string
	TopicNotif string
}

type MySQL struct {
	Host string
	Port string
	Usr  string
	Pw   string
	Db   string
}

// config provider
func NewConfig() *AppConfig {
	config := AppConfig{}

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}

	return &config
}
