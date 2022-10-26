package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/mudi25/pretest/errors"
)

func response(ctx echo.Context, data interface{}, err error) error {
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	code, msg := errors.FromError(err)
	type res struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `data:"data"`
	}
	return ctx.JSON(code, res{Code: code, Message: msg, Data: data})
}
