package http

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"lion/internal/interfaces/container"
	"os"
	"os/signal"
	"time"
)

func Start(container *container.Container) {
	server := echo.New()
	SetupMiddleware(server, container)
	SetupRoutes(container, server, NewHandler(container))

	port := fmt.Sprintf(":%v", container.Config.App.Port)
	go func() {
		if err := server.Start(port); err != nil {
			server.Logger.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		server.Logger.Fatal(err)
	}
}
