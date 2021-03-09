package service

import (
	"context"
	"gee/model"
)

// UserService ...
type UserService interface {
	RegisterUser(ctx context.Context, user *model.User) (ruser *model.User, err error)
}
