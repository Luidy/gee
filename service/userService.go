package service

import (
	"context"
	"gee/model"
	"gee/repository"
	"strconv"
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

func (u *userService) GetUser(ctx context.Context, id string) (ruser *model.User, err error) {
	uid, _ := strconv.ParseUint(id, 0, 64)
	return u.uRepo.GetUser(ctx, uid)
}

func (u *userService) UpdateUser(ctx context.Context, user *model.User) (ruser *model.User, err error) {
	ruser, err = u.uRepo.GetUser(ctx, uint64(user.ID))
	if err != nil {
		return nil, err
	}
	ruser.UpdateUser(user)
	if err := u.uRepo.UpdateUser(ctx, ruser); err != nil {
		return nil, err
	}
	return ruser, nil
}

func (u *userService) DeleteUser(ctx context.Context, id string) (ruser *model.User, err error) {
	uid, _ := strconv.ParseUint(id, 0, 64)
	ruser, err = u.uRepo.GetUser(ctx, uid)
	if err != nil {
		return nil, err
	}
	if err := u.uRepo.DeleteUser(ctx, uid); err != nil {
		return nil, err
	}
	return ruser, nil
}
