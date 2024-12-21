package main

import (
	"calc/internal/app"
	"context"
	"fmt"
	"log"
	"os"
	"syscall"

	exit "calc/pkg/context"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Ошибка чтения конфигурации: %s", err)
	}
}

func main() {
	app, err := app.New()
	if err != nil {
		fmt.Println(err)
		return
	}

	ctx, cancel := exit.WithSignal(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	if err := app.Run(ctx); err != nil {
		fmt.Println(err)
		return
	}

	return
}
