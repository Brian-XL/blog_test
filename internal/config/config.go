package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	App      App
	Database DatabaseConfig
	JWT      JWT
}

type App struct {
	Name string
	Port uint
}

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
	SSLMode  string
	TimeZone string
}

type JWT struct {
	Secret      string
	ExpireHours uint
}

func LoadConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	cfg := &Config{}

	if err := viper.Unmarshal(cfg); err != nil {
		log.Fatalf("Unable to decode config into struct: %v", err)
	}

	return cfg
}
