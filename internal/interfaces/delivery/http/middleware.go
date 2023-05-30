package http

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"lion/internal/interfaces/container"
	validate "lion/internal/shared/helper/validator"
	mware "lion/internal/shared/middleware"
	"net/http"
)

func SetupMiddleware(server *echo.Echo, container *container.Container) {
	server.HTTPErrorHandler = mware.EchoErrorHandler
	server.Use(middleware.Recover())
	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodOptions},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderAuthorization,
			echo.HeaderAccessControlAllowOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderContentLength,
			echo.HeaderAcceptEncoding,
			echo.HeaderXCSRFToken,
			echo.HeaderXRequestID,
		},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))
	server.HideBanner = true
	server.Validator = &validate.CustomValidator{Validator: validator.New()}
}
