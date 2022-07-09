package http

import "github.com/labstack/echo/v4"

func (uh *UserHandler) SetRoutes(e *echo.Echo) {
	e.GET("/user/:id", uh.GetOneById)
}
