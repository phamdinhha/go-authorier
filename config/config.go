package config

import (
	"errors"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig   `mapstructure:"SERVER"`
	Logger   Logger         `mapstructure:"LOGGER"`
	Postgres PostgresConfig `mapstructure:"POSTGRES"`
	Casbin   CasbinConfig   `mapstructure:"CASBIN"`
}

type ServerConfig struct {
	Host        string `mapstructure:"HOST"`
	Port        string `mapstructure:"PORT"`
	Environment string `mapstructure:"ENVIRONMENT"`
}

type Logger struct {
	Development bool   `mapstructure:"DEVELOPMENT"`
	Encoding    string `mapstructure:"ENCODING"`
	Level       string `mapstructure:"LEVEL"`
}

type PostgresConfig struct {
	Host     string `mapstructure:"HOST"`
	Port     string `mapstructure:"PORT"`
	User     string `mapstructure:"USER"`
	Password string `mapstructure:"PASSWORD"`
	Database string `mapstructure:"DATABASE"`
	Driver   string `mapstructure:"DRIVER"`
}

type CasbinConfig struct {
	DbMode        bool   `mapstructure:"MODE"`
	PostgresTable string `mapstructure:"POSTGRES_TABLE"`
	ModelConfig   string `mapstructure:"MODEL_CONFIG"`
	Policy        string `mapstructure:"POLICY"`
}

func LoadConfig(filename string) (*viper.Viper, error) {
	v := viper.New()

	v.SetConfigName(filename)
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
