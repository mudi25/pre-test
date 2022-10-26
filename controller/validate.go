package controller

import (
	"github.com/labstack/echo/v4"
)

func validateRequest(ctx echo.Context, data interface{}) error {
	var err error
	if err = ctx.Bind(data); err == nil {
		if err = validate.Struct(data); err == nil {
			return nil
		}
	}
	return err
}
