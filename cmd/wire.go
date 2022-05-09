//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"
	"time"

	"github.com/CO88/go-ddd-boilerplate/api"
	"github.com/CO88/go-ddd-boilerplate/user/delivery/http"
	"github.com/CO88/go-ddd-boilerplate/user/repository/mysql"
	"github.com/CO88/go-ddd-boilerplate/user/usecase"
	"github.com/google/wire"
)

// var Mainset = wire.NewSet(http.NewUserHandler, config.LoadConfig, container.NewDIContainer)

// func InitailizeHandler(db *sql.DB) *container.DIContainer {
// 	wire.Build(Mainset)
// 	return nil
// }

func InitailizeHandler(dbConn *sql.DB, timeout time.Duration) api.ApiHandler {
	panic(wire.Build(
		mysql.NewUserRepository,
		usecase.NewUserUsecase,
		http.NewUserHandler,
	))
}
