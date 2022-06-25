package config

import (
	"errors"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Logger   Logger
	Postgres PostgresConfig
}

type ServerConfig struct {
	Host        string
	Port        string
	Environment string
}

type Logger struct {
	Development bool
	Encoding    string
	Level       string
}

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func LoadConfig(fileName string) (*viper.Viper, error) {
	v := viper.New()

	v.SetConfigName(fileName)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}
	return v, nil
}

func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config
	err := v.Unmarshal(&c)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}
	return &c, nil
}
