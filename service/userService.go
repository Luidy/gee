package service

import (
	"context"
	"gee/model"
	"gee/repository"
)

type userService struct {
	uRepo repository.UserRepository
}

// NewUserService ...
func NewUserService(uRepo repository.UserRepository) UserService {
	return &userService{
		uRepo: uRepo,
	}
}

func (u *userService) RegisterUser(ctx context.Context, user *model.User) (ruser *model.User, err error) {
	return u.uRepo.CreateUser(ctx, user)
}
