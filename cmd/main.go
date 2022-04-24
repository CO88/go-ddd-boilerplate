package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/CO88/go-ddd-boilerplate/config"
	_userHandler "github.com/CO88/go-ddd-boilerplate/user/delivery/http"
	"github.com/CO88/go-ddd-boilerplate/user/repository/mysql"
	"github.com/CO88/go-ddd-boilerplate/user/usecase"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	configuration := config.LoadConfig()

	connection := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		configuration.Db.UserName,
		configuration.Db.UserPassword,
		configuration.Db.Host,
		configuration.Db.Port,
		configuration.Db.Name)

	dbConn, err := sql.Open(`mysql`, connection)
	if err != nil {
		log.Fatal(err)
	}

	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	e := echo.New()
	e.Use(middleware.Recover())

	//
	userRepo := mysql.NewUserRepository(dbConn)

	timeoutContext := time.Duration(configuration.Context.Timeout) * time.Second
	uu := usecase.NewUserUsecase(userRepo, timeoutContext)
	_userHandler.NewUserHandler(e, uu)
	//

	fmt.Println("hi")
	fmt.Println(configuration)

	log.Fatal(e.Start(configuration.Server.Address))
}
