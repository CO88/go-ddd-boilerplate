package domain

import (
	"context"
	"time"
)

type User struct {
	ID          int64     `json:"id"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Name        string    `json:"name"`
	PhoneNumber string    `json:"phone_number"`
	CreateAt    time.Time `json:"created_at"`
	UpdateAt    time.Time `json:"updated_at"`
}

type UserUsecase interface {
	GetOneById(ctx context.Context, id int64) (User, error)
}

type UserRespository interface {
	FindOneById(ctx context.Context, id int64) (User, error)
}
