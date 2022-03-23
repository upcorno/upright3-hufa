package controller

import (
	"law/enum"
	"law/model"
	"law/utils"
	"time"

	"github.com/labstack/echo/v4"
)

//创建咨询实例
func ConsultationCreate(ctx echo.Context) error {
	consul := &model.Consultation{}
	if err := ctx.Bind(consul); err != nil {
		return ctx.JSON(utils.ErrIpt("输入解析失败！", err.Error()))
	}
	uid := ctx.Get("uid").(int)
	consul.ConsultantUid = uid
	consul.Status = enum.DOING
	consul.CreateTime = int(time.Now().Unix())
	if err := ctx.Validate(consul); err != nil {
		return ctx.JSON(utils.ErrIpt("输入校验失败！", err.Error()))
	}
	if err := model.ConsultationCreate(consul); err != nil {
		return ctx.JSON(utils.ErrIpt("法律咨询生成失败！", err.Error()))
	}
	return ctx.JSON(utils.Succ("success", map[string]int{"consultation_id": consul.Id}))
}