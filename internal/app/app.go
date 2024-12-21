package app

import (
	"calc/config"
	"calc/internal/handler"
	"calc/internal/service"
	"calc/pkg/calc"
	"context"
	"errors"
	"fmt"
	"net/http"
)

type App struct {
	server *http.Server
}

func New() (*App, error) {
	cfg, err := config.Init()
	if err != nil {
		return nil, err
	}

	calculator := calc.New()

	srv := service.New(calculator)

	result := &App{
		server: &http.Server{
			Addr:    cfg.Srv(),
			Handler: handler.New(srv),
		},
	}

	return result, nil
}

func (a *App) Run(ctx context.Context) error {
	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("Shutting down the server...")

			err := a.server.Shutdown(context.Background())
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println("Server shutting down successfully")

			return
		}
	}()

	fmt.Printf("Server is running on http://localhost%s \n", a.server.Addr)

	if err := a.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	fmt.Println("Server stopped")

	return nil
}
