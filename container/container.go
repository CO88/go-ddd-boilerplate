package container

import (
	"database/sql"
	"time"

	"github.com/CO88/go-ddd-boilerplate/api"
	"github.com/CO88/go-ddd-boilerplate/user/delivery/http"
	"github.com/CO88/go-ddd-boilerplate/user/repository/mysql"
	"github.com/CO88/go-ddd-boilerplate/user/usecase"
)

// type DIContainer struct {
// 	Config   *config.Configuration
// 	Handlers []api.ApiHandler
// }

// var Container *DIContainer

// func NewContainer(dbConn *sql.DB, timeoutContext time.Duration) api.ApiHandler {
// 	panic(wire.Build(
// 		mysql.NewUserRepository,
// 		usecase.NewUserUsecase,
// 		http.NewUserHandler,
// 	))
// }

func NewContainer(dbConn *sql.DB, timeout time.Duration) api.ApiHandler {

	userRepo := mysql.NewUserRepository(dbConn)
	userUsecase := usecase.NewUserUsecase(userRepo, timeout)
	userHandler := http.NewUserHandler(userUsecase)

	return userHandler
}
