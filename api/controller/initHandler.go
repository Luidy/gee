package controller

import (
	"gee/config"
	"gee/repository"
	"gee/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// GeeResponse ...
type GeeResponse struct {
	ResultMsg  string
	ResultData interface{}
}

// InitHandler ...
func InitHandler(gee *config.ViperConfig, e *echo.Echo, db *gorm.DB) (err error) {
	api := e.Group("/api")
	version := api.Group("/v1")
	eg := version.Group("/gee")

	uRepo := repository.NewGormUserRepository(db)
	uSvc := service.NewUserService(uRepo)
	newHTTPHandler(gee, eg, uSvc)
	return nil
}

func response(c echo.Context, code int, msg string, result ...interface{}) (err error) {
	res := GeeResponse{
		ResultMsg:  msg,
		ResultData: result,
	}
	return c.JSON(code, res)
}
