package usecase

import (
	"context"
	"time"

	"github.com/CO88/go-ddd-boilerplate/user/domain"
)

type userUsecase struct {
	userRepository domain.UserRespository
	contextTimeout time.Duration
}

func NewUserUsecase(uu domain.UserRespository, timeout time.Duration) domain.UserUsecase {
	return &userUsecase{
		userRepository: uu,
		contextTimeout: timeout,
	}
}

func (u *userUsecase) GetOneById(ctx context.Context, id int64) (res domain.User, err error) {
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	res, err = u.userRepository.FindOneById(ctx, id)
	if err != nil {
		return
	}

	return
}
