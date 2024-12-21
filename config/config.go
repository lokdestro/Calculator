package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int    `mapstructure:"port"`
		Host string `mapstructure:"host"`
	} `mapstructure:"server"`
}

func (c Config) Srv() string {
	return fmt.Sprintf(":%d", c.Server.Port)
}

func Init() (Config, error) {
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Ошибка декодирования конфигурации: %s", err)
	}

	return config, nil
}
