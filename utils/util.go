package utils

import "github.com/labstack/echo/v4"

func BindAndValidate(ctx echo.Context, i interface{}) (err error) {
	if err = ctx.Bind(i); err == nil {
		err = ctx.Validate(i)
	}
	return
}
