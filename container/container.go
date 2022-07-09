package container

import (
	"github.com/CO88/go-ddd-boilerplate/api"
	"github.com/CO88/go-ddd-boilerplate/config"
	user "github.com/CO88/go-ddd-boilerplate/user/delivery/http"
)

type DIContainer struct {
	Config   *config.Configuration
	Handlers []api.ApiHandler
}

var Container *DIContainer

func NewDIContainer(
	configuration *config.Configuration,
	userHanlder *user.UserHandler,
) *DIContainer {
	Container := &DIContainer{
		configuration,
		[]api.ApiHandler{userHanlder},
	}

	return Container
}
