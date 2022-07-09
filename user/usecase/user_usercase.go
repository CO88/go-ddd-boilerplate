package usecase

import (
	"context"
	"time"

	"github.com/CO88/go-ddd-boilerplate/config"
	"github.com/CO88/go-ddd-boilerplate/user/domain"
)

type userUsecase struct {
	userRepository domain.UserRespository
	cfg            *config.Configuration
}

func NewUserUsecase(cfg *config.Configuration, uu domain.UserRespository) domain.UserUsecase {
	return &userUsecase{
		userRepository: uu,
		cfg:            cfg,
	}
}

func (u *userUsecase) GetOneById(ctx context.Context, id int64) (res domain.User, err error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(u.cfg.Context.Timeout))
	defer cancel()

	res, err = u.userRepository.FindOneById(ctx, id)
	if err != nil {
		return
	}

	return
}
