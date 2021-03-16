package controller

import (
	"context"
	"gee/config"
	"gee/model"
	"gee/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type httpHandler struct {
	uSvc service.UserService
}

func newHTTPHandler(gee *config.ViperConfig,
	eg *echo.Group,
	uSvc service.UserService) {

	handler := &httpHandler{
		uSvc: uSvc,
	}
	eg.POST("/register", handler.RegisterUser)

	userGroup := eg.Group("/user")
	newHTTPUserHandler(userGroup, handler)
}

func (h *httpHandler) RegisterUser(c echo.Context) (err error) {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	user := &model.User{}
	if err := c.Bind(user); err != nil {
		return response(c, http.StatusBadRequest, err.Error())
	}

	if user, err = h.uSvc.RegisterUser(ctx, user); err != nil {
		return response(c, http.StatusInternalServerError, err.Error())
	}

	return response(c, http.StatusOK, "RegisterUser", user)
}
