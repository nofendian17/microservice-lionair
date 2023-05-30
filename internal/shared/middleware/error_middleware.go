package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"lion/internal/model/response"
)

// EchoErrorHandler handles errors in Echo middleware
func EchoErrorHandler(err error, c echo.Context) {
	if c.Response().Committed || err == nil {
		return
	}

	switch {
	case errors.Is(err, echo.ErrNotFound):
		c.JSON(http.StatusNotFound, response.Response{
			Code:    http.StatusNotFound,
			Message: http.StatusText(http.StatusNotFound),
		})
	case errors.Is(err, echo.ErrInternalServerError):
		c.JSON(http.StatusInternalServerError, response.Response{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
		})
	case errors.As(err, new(validator.ValidationErrors)):
		validationErrors := err.(validator.ValidationErrors)
		var messages []string
		for _, v := range validationErrors {
			message := fmt.Sprintf("%s: invalid value '%v'", v.Field(), v.Value())
			messages = append(messages, message)
		}
		c.JSON(http.StatusBadRequest, response.Response{
			Code:    http.StatusBadRequest,
			Message: strings.Join(messages, "; "),
		})
	default:
		c.JSON(http.StatusInternalServerError, response.Response{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
		})
	}
}
