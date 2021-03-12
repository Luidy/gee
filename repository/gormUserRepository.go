package repository

import (
	"context"
	"gee/model"
	"log"

	"github.com/juju/errors"
	"gorm.io/gorm"
)

type gormUserRepository struct {
	Conn *gorm.DB
}

// NewGormUserRepository ...
func NewGormUserRepository(conn *gorm.DB) UserRepository {
	if err := conn.AutoMigrate(&model.User{}); err != nil {
		log.Panic("Unable to Automigrate UserReository.")
	}

	return &gormUserRepository{Conn: conn}
}

func (g *gormUserRepository) CreateUser(ctx context.Context, user *model.User) (ruser *model.User, err error) {
	scope := g.Conn.WithContext(ctx)
	if err := scope.Create(&user).Error; err != nil {
		log.Println("gormUserRepository RegisterUser err", err)
		return nil, err
	}

	return user, nil
}

func (g *gormUserRepository) GetUserList(ctx context.Context) (rusers model.Users, err error) {
	scope := g.Conn.WithContext(ctx)
	if scope.Find(&rusers).RowsAffected == 0 {
		log.Println("gormUserRepository GetUserList Not Found.")
	}
	return rusers, nil
}

func (g *gormUserRepository) GetUser(ctx context.Context, id uint64) (ruser *model.User, err error) {
	scope := g.Conn.WithContext(ctx)
	if scope.Where("id = ?", id).Find(&ruser).RowsAffected == 0 {
		return nil, errors.NotFoundf("gormUserRepository GetUser Not Found. id[%v]", id)
	}
	return ruser, nil
}

func (g *gormUserRepository) UpdateUser(ctx context.Context, id uint64) (err error) {
	scope := g.Conn.WithContext(ctx)
	if err := scope.Updates(&user).Error; err != nil {
		log.Println("gormUserRepository UpdateUser err", err)
		return err
	}

	return nil
}

func (g *gormUserRepository) DeleteUser(ctx context.Context, id uint64) (err error) {
	scope := g.Conn.WithContext(ctx)
	if err := scope.Delete(&user).Error; err != nil {
		log.Println("gormUserRepository DeleteUser err", err)
		return err
	}

	return nil
}
