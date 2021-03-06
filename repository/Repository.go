package repository

import (
	"context"
	"fmt"
	"gee/config"
	"gee/model"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// InitDB ...
func InitDB(gee *config.ViperConfig) *gorm.DB {
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&loc=UTC",
		gee.GetString("db_user"),
		gee.GetString("db_pass"),
		gee.GetString("db_host"),
		gee.GetInt("db_port"),
		gee.GetString("db_name"))

	dbConn, err := gorm.Open(mysql.Open(dbURI))
	if err != nil {
		log.Println("InitDB", "err", err)
	}
	if dbConn == nil {
		os.Exit(1)
	}
	return dbConn
}

// UserRepository ...
type UserRepository interface {
	CreateUser(ctx context.Context, user *model.User) (ruser *model.User, err error)
	GetUserList(ctx context.Context) (rusers model.Users, err error)
	GetUser(ctx context.Context, id uint64) (ruser *model.User, err error)
	UpdateUser(ctx context.Context, user *model.User) (err error)
	DeleteUser(ctx context.Context, id uint64) (err error)
}
