package controller

import (
	"law/model"
	"law/utils"
	"strconv"

	"github.com/labstack/echo/v4"
)

//编辑回访记录
func ProtectionReturnVisitUpdate(ctx echo.Context) error {
	protectionIdStr := ctx.QueryParam("protection_id")
	protectionId, err := strconv.Atoi(protectionIdStr)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("获取protection_id失败！", err.Error()))
	}
	returnVisit := &model.ProtectionReturnVisit{}
	if err := ctx.Bind(returnVisit); err != nil {
		return ctx.JSON(utils.ErrIpt("输入解析失败！", err.Error()))
	}
	returnVisit.ProtectionId = protectionId
	if err := ctx.Validate(returnVisit); err != nil {
		return ctx.JSON(utils.ErrIpt("输入校验失败！", err.Error()))
	}
	if err := model.ProtectionRetureVisitUpdate(returnVisit); err != nil {
		return ctx.JSON(utils.ErrIpt("更新回访记录失败！", err.Error()))
	}
	return ctx.JSON(utils.Succ("success"))
}

//获取回访记录
func ProtectionReturnVisitGet(ctx echo.Context) error {
	protectionIdStr := ctx.QueryParam("protection_id")
	protectionId, err := strconv.Atoi(protectionIdStr)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("获取detection_id失败！", err.Error()))
	}
	returnVisit, err := model.ProtectionReturnVisitGet(protectionId)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("获取维权意向回访记录失败！", err.Error()))
	}
	return ctx.JSON(utils.Succ("success", returnVisit))
}
