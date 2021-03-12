package service

import (
	"context"
	"gee/model"
)

// UserService ...
type UserService interface {
	RegisterUser(ctx context.Context, user *model.User) (ruser *model.User, err error)
	GetUserList(ctx context.Context) (rusers model.Users, err error)
	GetUser(ctx context.Context, id uint64) (ruser *model.User, err error)
	UpdateUser(ctx context.Context, user *model.User) (ruser *model.User, err error)
	DeleteUser(ctx context.Context, id uint64) (ruser *model.User, err error)
}
