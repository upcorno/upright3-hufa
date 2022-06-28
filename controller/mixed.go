package controller

import (
	"law/service"
	"law/utils"

	zlog "github.com/rs/zerolog/log"

	"github.com/labstack/echo/v4"
)

type mixed struct{}

var MixedController *mixed

func (c *mixed) GetFundUsage(ctx echo.Context) error {
	obligeeProvince := ctx.QueryParam("obligee_province")
	if obligeeProvince == "" {
		return ctx.JSON(utils.ErrIpt("必须指定权利人所在省"))
	}
	usage, err := service.MixedSvr.GetFundUsage(obligeeProvince)
	if err != nil {
		zlog.Warn().Err(err).Msg(ctx.Path() + "执行失败")
		return ctx.JSON(utils.ErrSvr("查询失败"))
	}
	return ctx.JSON(utils.Succ("success", usage))
}
