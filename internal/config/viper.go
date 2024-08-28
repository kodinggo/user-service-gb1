package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	JWT JWT `yaml:"jwt"`
}

type JWT struct {
	Secret string        `yaml:"secret"`
	TTL    time.Duration `yaml:"ttl"`
}

var Cfg Config

func LoadWithViper() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".") // Look for config.yaml in the current directory

	// Read the configuration file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&Cfg)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}
}
