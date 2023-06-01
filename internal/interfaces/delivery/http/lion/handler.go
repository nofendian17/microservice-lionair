package lion

import (
	"context"
	"github.com/labstack/echo/v4"
	model "lion/internal/model/lion"
	"net/http"
	"time"
)

func (h *Handler) Schedule(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(h.config.App.HttpRequestTimeout)*time.Second)
	defer cancel()

	var request model.ScheduleRequest
	if err := c.Bind(&request); err != nil {
		return err
	}

	if err := c.Validate(&request); err != nil {
		return err
	}

	schedule, err := h.lionUseCase.ScheduleWithoutSession(ctx, request)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, schedule)
}

func (h *Handler) Pricing(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(h.config.App.HttpRequestTimeout)*time.Second)
	defer cancel()

	var request model.PricingRequest
	if err := c.Bind(&request); err != nil {
		return err
	}

	if err := c.Validate(&request); err != nil {
		return err
	}

	schedule, err := h.lionUseCase.Pricing(ctx, request)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, schedule)
}
