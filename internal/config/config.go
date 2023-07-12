package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env        string `yaml:"env" env:"ENV" end-default:"local"`
	BotToken   string `yaml:"bottoken" env-required:"true"`
	WebhookURL string `yaml:"webhookurl" env-required:"true"`
	Port       string `yaml:"port"`
}

func MustLoad() *Config {
	path := "news-bot/config/config.yaml"
	configPath := os.Getenv("ROOT_PATH") + "/" + path
	if configPath == "" {
		log.Fatal("Config path is not set")
	}

	//check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}
