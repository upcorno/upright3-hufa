package controller

import (
	"law/enum"
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
	err = model.ProtectionRetureVisitAdd(&model.ProtectionReturnVisit{
		CreatorUid: uid,
		ProtectionId: rightsProtection.Id,
		Classification: enum.NORETURN,
		CreateTime: int(time.Now().Unix()),
	})
	if err != nil {
		return ctx.JSON(utils.ErrIpt("添加默认回访失败！", err.Error()))
	}
	return ctx.JSON(utils.Succ("success", map[string]int{"rights_protection_id": rightsProtection.Id}))
}

//获取维权意向
func RightsProtectionGet(ctx echo.Context) error {
	protectionIdStr := ctx.QueryParam("protection_id")
	protectionId, err := strconv.Atoi(protectionIdStr)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("获取protection_id失败！", err.Error()))
	}
	rightsProtection, err := model.RightsProtectionGet(protectionId)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("获取维权意向失败！", err.Error()))
	}
	return ctx.JSON(utils.Succ("success", rightsProtection))
}

//侵权监测列表检索
func RightsProtectionList(ctx echo.Context) error {
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
	protections, err := service.RightsProtectionList(page, search)
	if err != nil {
		return ctx.JSON(utils.ErrSvr("获取rights_protection list失败", err.Error()))
	}
	return ctx.JSON(utils.Succ("success", protections))
}