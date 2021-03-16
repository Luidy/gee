package controller

import (
	"context"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func newHTTPUserHandler(eg *echo.Group, handler *httpHandler) {
	// prefix: /api/v1/gee/user
	eg.GET("", handler.GetUserList)
	eg.GET("/:id", handler.GetUser)
	eg.DELETE("/:id", handler.DeleteUser)
}

func (h *httpHandler) GetUserList(c echo.Context) (err error) {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	users, err := h.uSvc.GetUserList(ctx)
	if err != nil {
		log.Println("GetUserList err", err)
		return response(c, http.StatusInternalServerError, err.Error())
	}
	return response(c, http.StatusOK, "GetUserList OK", users)
}

func (h *httpHandler) GetUser(c echo.Context) (err error) {
	id := c.Param("id")
	if id == "" {
		return response(c, http.StatusBadRequest, "User ID should not be empty.")
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	user, err := h.uSvc.GetUser(ctx, id)
	if err != nil {
		log.Println("GetUser err", err)
		return response(c, http.StatusInternalServerError, err.Error())
	}
	return response(c, http.StatusOK, "GetUser OK", user)
}

func (h *httpHandler) DeleteUser(c echo.Context) (err error) {
	id := c.Param("id")
	if id == "" {
		return response(c, http.StatusBadRequest, "User ID should not be empty.")
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	user, err := h.uSvc.DeleteUser(ctx, id)
	if err != nil {
		log.Println("DeleteUser err", err)
		return response(c, http.StatusInternalServerError, err.Error())
	}
	return response(c, http.StatusOK, "DeleteUser OK", user)
}
