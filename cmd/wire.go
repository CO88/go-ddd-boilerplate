//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/CO88/go-ddd-boilerplate/config"
	"github.com/CO88/go-ddd-boilerplate/container"
	"github.com/CO88/go-ddd-boilerplate/user"
	"github.com/google/wire"
)

var MainSet = wire.NewSet(container.NewDIContainer, user.UserSet)

func InitailizeDIContainer(configuration *config.Configuration, conn *sql.DB) *container.DIContainer {
	wire.Build(MainSet)
	return nil
}
