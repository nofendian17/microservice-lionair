package http

import (
	"github.com/labstack/echo/v4"
	"lion/internal/interfaces/container"
)

const (
	VERSION1 = "/api/v1"
)

func SetupRoutes(container *container.Container, e *echo.Echo, handler *Handler) {
	// Health routes
	e.GET("/", handler.health.Hello)
	e.GET("/health", handler.health.Health)
	e.GET("/ping", handler.health.Ping)
	e.GET("/ready", handler.health.Readiness)

	// V1 routes
	v1 := e.Group(VERSION1)
	v1.POST("/schedule", handler.lion.Schedule)
	v1.POST("/pricing", handler.lion.Pricing)
}
