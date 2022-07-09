package main

import (
	"log"

	"github.com/CO88/go-ddd-boilerplate/config"
	"github.com/CO88/go-ddd-boilerplate/tools/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	configuration := config.LoadConfig()

	conn := mysql.NewMysqlConn(*configuration)

	e := echo.New()
	e.Use(middleware.Recover())

	container := InitailizeDIContainer(configuration, conn)

	for _, h := range container.Handlers {
		h.SetRoutes(e)
	}

	log.Fatal(e.Start(configuration.Server.Address))
}
