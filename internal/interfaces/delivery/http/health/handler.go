package health

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"lion/internal/model/health"
	"net/http"
)

func (h *Handler) Hello(c echo.Context) error {
	return c.String(http.StatusOK, fmt.Sprintf("%s is running", h.config.App.Name))
}

func (h *Handler) Ping(c echo.Context) (err error) {
	return c.String(http.StatusOK, "pong")
}

func (h *Handler) Readiness(c echo.Context) error {
	return c.JSON(http.StatusOK, h.healthUseCase.Readiness(context.Background()))
}

func (h *Handler) Health(c echo.Context) error {
	response := health.Readiness{
		Status: http.StatusText(http.StatusOK),
	}
	return c.JSON(http.StatusOK, response)
}
