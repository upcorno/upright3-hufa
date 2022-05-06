package controller

import (
	"law/model"
	"law/service"
	"law/utils"
	"strconv"

	"github.com/labstack/echo/v4"
)

//添加侵权监测
func InfringementMonitorAdd(ctx echo.Context) error {
	baseInfo := &service.InfringementMonitorBaseInfo{}
	if err := utils.BindAndValidate(ctx, baseInfo); err != nil {
		return ctx.JSON(utils.ErrIpt("输入解析校验失败！", err.Error()))
	}
	uid := ctx.Get("uid").(int)
	beanId, err := service.InfringementMonitorAdd(baseInfo, uid)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("添加侵权监测失败！", err.Error()))
	}
	return ctx.JSON(utils.Succ("success", map[string]int{"infringement_monitor_id": beanId}))
}

func InfringementMonitorGet(ctx echo.Context) error {
	uid := ctx.Get("uid").(int)
	bean, err := service.InfringementMonitorGet(uid)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("获取侵权监测失败！", err.Error()))
	}
	return ctx.JSON(utils.Succ("success", bean))
}

func InfringementMonitorBgGet(ctx echo.Context) error {
	beanIdStr := ctx.QueryParam("id")
	beanId, err := strconv.Atoi(beanIdStr)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("获取id失败", err.Error()))
	}
	bean, err := service.InfringementMonitorBgGet(beanId)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("获取侵权监测失败！", err.Error()))
	}
	return ctx.JSON(utils.Succ("success", bean))
}

func InfringementMonitorSetDealInfo(ctx echo.Context) error {
	dealInfo := &service.InfringementMonitorDealInfo{}
	beanIdStr := ctx.QueryParam("id")
	beanId, err := strconv.Atoi(beanIdStr)
	if err == nil {
		dealInfo.Id = beanId
	}
	if err := utils.BindAndValidate(ctx, dealInfo); err != nil {
		return ctx.JSON(utils.ErrIpt("输入解析校验失败！", err.Error()))
	}
	if err := service.InfringementMonitorSetDealInfo(dealInfo.Id, dealInfo); err != nil {
		return ctx.JSON(utils.ErrIpt("设置回访记录失败！", err.Error()))
	}
	return ctx.JSON(utils.Succ("success"))
}

func InfringementMonitorUpdateBaseInfo(ctx echo.Context) error {
	baseInfo := &service.InfringementMonitorBaseInfo{}
	if err := utils.BindAndValidate(ctx, baseInfo); err != nil {
		return ctx.JSON(utils.ErrIpt("输入解析校验失败！", err.Error()))
	}
	uid := ctx.Get("uid").(int)
	if err := service.InfringementMonitorUpdateBaseInfo(uid, baseInfo); err != nil {
		return ctx.JSON(utils.ErrIpt("修改基础信息失败！", err.Error()))
	}
	return ctx.JSON(utils.Succ("success"))
}

//侵权监测列表检索
func InfringementMonitorBackendList(ctx echo.Context) error {
	page := &model.Page{PageIndex: 1, ItemNum: 10}
	if err := ctx.Bind(page); err != nil {
		return ctx.JSON(utils.ErrIpt("分页输入错误,请重试！", err.Error()))
	}
	if err := ctx.Validate(page); err != nil {
		return ctx.JSON(utils.ErrIpt("分页数据输入校验失败！", err.Error()))
	}
	search := &service.InfringementMonitorSearch{}
	if err := ctx.Bind(search); err != nil {
		return ctx.JSON(utils.ErrIpt("检索数据输入错误,请重试！", err.Error()))
	}
	if err := ctx.Validate(search); err != nil {
		return ctx.JSON(utils.ErrIpt("检索输入校验失败！", err.Error()))
	}
	beans, err := service.InfringementMonitorBackendList(page, search)
	if err != nil {
		return ctx.JSON(utils.ErrSvr("获取infringement_monitor list失败", err.Error()))
	}
	return ctx.JSON(utils.Succ("success", beans))
}
