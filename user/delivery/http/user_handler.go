package http

import (
	"net/http"
	"strconv"

	"github.com/CO88/go-ddd-boilerplate/api"
	"github.com/CO88/go-ddd-boilerplate/user/domain"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type ResponseError struct {
	Message string `json:"message"`
}

type UserHandler struct {
	UserUsecase domain.UserUsecase
}

func NewUserHandler(uu domain.UserUsecase) api.ApiHandler {
	return &UserHandler{UserUsecase: uu}
}

func (uh *UserHandler) SetRoutes(e *echo.Echo) {
	e.GET("/user/:id", uh.GetOneById)
}

func (uh *UserHandler) GetOneById(c echo.Context) error {
	idp, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}

	id := int64(idp)
	ctx := c.Request().Context()

	user, err := uh.UserUsecase.GetOneById(ctx, id)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, user)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
