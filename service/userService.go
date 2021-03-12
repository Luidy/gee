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

func (u *userService) GetUserList(ctx context.Context) (rusers model.Users, err error) {
	return u.uRepo.GetUserList(ctx)
}

func (u *userService) GetUser(ctx context.Context, id uint64) (ruser *model.User, err error) {
	return u.uRepo.GetUser(ctx, id)
}

func (u *userService) UpdateUser(ctx context.Context, user *model.User) (ruser *model.User, err error) {
	if ruser, err := u.uRepo.GetUser(ctx, user.ID); err != nil {
		return nil, err
	}
	ruser = ruser.UpdateUser(user)
	if err := u.uRepo.UpdateUser(ctx, ruser); err != nil {
		return nil, err
	}
	return ruser, nil
}

func (u *userService) DeleteUser(ctx context.Context, id uint64) (ruser *model.User, err error) {
	if ruser, err := u.uRepo.GetUser(ctx, id); err != nil {
		return nil, err
	}
	if err := u.uRepo.DeleteUser(ctx, id); err != nil {
		return nil, err
	}
	return ruser, nil
}
