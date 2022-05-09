package api

import "github.com/labstack/echo/v4"

type ApiHandler interface {
	SetRoutes(e *echo.Echo)
}
