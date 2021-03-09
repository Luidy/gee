package repository

import (
	"context"
	"gee/model"
	"log"

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
