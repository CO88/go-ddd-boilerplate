package user

import (
	user "github.com/CO88/go-ddd-boilerplate/user/delivery/http"
	"github.com/CO88/go-ddd-boilerplate/user/repository"
	"github.com/CO88/go-ddd-boilerplate/user/usecase"
	"github.com/google/wire"
)

var UserSet = wire.NewSet(
	user.NewUserHandler,
	usecase.NewUserUsecase,
	repository.NewUserRepository,
)
