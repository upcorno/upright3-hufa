package controller

import (
	"law/model"
	"law/service"
	"law/utils"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

//添加维权意向
func RightsProtectionAdd(ctx echo.Context) error {
	rightsProtection := &model.RightsProtection{}
	if err := ctx.Bind(rightsProtection); err != nil {
		return ctx.JSON(utils.ErrIpt("输入解析失败！", err.Error()))
	}
	uid := ctx.Get("uid").(int)
	rightsProtection.CreatorUid = uid
	rightsProtection.CreateTime = int(time.Now().Unix())
	err := model.RightsProtectionAdd(rightsProtection)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("添加维权意向失败！", err.Error()))
	}
	return ctx.JSON(utils.Succ("success", map[string]int{"rights_protection_id": rightsProtection.Id}))
}

//获取维权意向
func RightsProtectionGet(ctx echo.Context) error {
	protectionIdStr := ctx.QueryParam("id")
	protectionId, err := strconv.Atoi(protectionIdStr)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("获取id失败！", err.Error()))
	}
	protection, err := model.RightsProtectionGet(protectionId)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("获取维权意向失败！", err.Error()))
	}
	return ctx.JSON(utils.Succ("success", protection))
}
func RightsProtectionSetDealInfo(ctx echo.Context) error {
	protectionIdStr := ctx.QueryParam("id")
	protectionId, err := strconv.Atoi(protectionIdStr)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("获取id失败！", err.Error()))
	}
	dealInfo := &service.RightsProtectionDealInfo{Id: protectionId}
	if err := ctx.Bind(dealInfo); err != nil {
		return ctx.JSON(utils.ErrIpt("输入解析失败！", err.Error()))
	}
	if err := ctx.Validate(dealInfo); err != nil {
		return ctx.JSON(utils.ErrIpt("输入校验失败！", err.Error()))
	}
	if err := service.RightsProtectionSetDealInfo(dealInfo); err != nil {
		return ctx.JSON(utils.ErrIpt("设置回访记录失败！", err.Error()))
	}
	return ctx.JSON(utils.Succ("success"))
}

//侵权监测列表检索
func RightsProtectionBackendList(ctx echo.Context) error {
	page := &model.Page{PageIndex: 1, ItemNum: 10}
	if err := ctx.Bind(page); err != nil {
		return ctx.JSON(utils.ErrIpt("分页输入错误,请重试！", err.Error()))
	}
	if err := ctx.Validate(page); err != nil {
		return ctx.JSON(utils.ErrIpt("分页数据输入校验失败！", err.Error()))
	}
	search := &service.RightsProtectionSearch{}
	if err := ctx.Bind(search); err != nil {
		return ctx.JSON(utils.ErrIpt("检索数据输入错误,请重试！", err.Error()))
	}
	if err := ctx.Validate(search); err != nil {
		return ctx.JSON(utils.ErrIpt("检索输入校验失败！", err.Error()))
	}
	protections, err := service.RightsProtectionBackendList(page, search)
	if err != nil {
		return ctx.JSON(utils.ErrSvr("获取rights_protection list失败", err.Error()))
	}
	return ctx.JSON(utils.Succ("success", protections))
}
