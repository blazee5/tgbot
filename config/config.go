package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

type Config struct {
	Bot   `yaml:"bot"`
	DB    `yaml:"db"`
	Redis `yaml:"redis"`
}

type Bot struct {
	Token   string `yaml:"token" env:"BOT_TOKEN"`
	AdminID int    `yaml:"admin_id" ENV:"ADMIN_ID"`
}

type DB struct {
	Host     string `yaml:"host" env:"DB_HOST" env-default:"localhost"`
	Port     string `yaml:"port" env:"DB_PORT" env-default:"5432"`
	Name     string `yaml:"name" env:"DB_NAME" env-default:"tgbot"`
	User     string `yaml:"user" env:"DB_USER" env-default:"postgres"`
	Password string `yaml:"password" env:"DB_PASSWORD" env-default:""`
	SSLMode  string `yaml:"ssl_mode" env:"DB_SSL" env-default:"false"`
}

type Redis struct {
	Host     string `yaml:"host" env:"REDIS_HOST" env-default:"localhost"`
	Port     string `yaml:"port" env:"REDIS_PORT" env-default:"6379"`
	Password string `yaml:"password" env:"REDIS_PASSWORD" env-default:""`
}

func LoadConfig() *Config {
	var cfg Config

	err := cleanenv.ReadConfig("config/config.yml", &cfg)

	if err != nil {
		log.Fatalf("error while read config: %v", err)
	}

	return &cfg
}
