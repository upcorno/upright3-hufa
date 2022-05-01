package controller

import (
	"law/model"
	"law/service"
	"law/utils"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

//添加侵权监测
func InfringementMonitorAdd(ctx echo.Context) error {
	infringementMonitor := &model.InfringementMonitor{}
	if err := ctx.Bind(infringementMonitor); err != nil {
		return ctx.JSON(utils.ErrIpt("输入解析失败！", err.Error()))
	}
	if err := ctx.Validate(infringementMonitor); err != nil {
		return ctx.JSON(utils.ErrIpt("输入校验失败！", err.Error()))
	}
	uid := ctx.Get("uid").(int)
	infringementMonitor.CreatorUid = uid
	infringementMonitor.CreateTime = int(time.Now().Unix())
	err := model.InfringementMonitorAdd(infringementMonitor)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("添加侵权监测失败！", err.Error()))
	}
	return ctx.JSON(utils.Succ("success", map[string]int{"infringement_monitor_id": infringementMonitor.Id}))
}

//获取侵权监测
func InfringementMonitorGet(ctx echo.Context) error {
	monitorIdStr := ctx.QueryParam("id")
	monitorId, err := strconv.Atoi(monitorIdStr)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("获取id失败！", err.Error()))
	}
	monitor, err := model.InfringementMonitorGet(monitorId)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("获取侵权监测失败！", err.Error()))
	}
	return ctx.JSON(utils.Succ("success", monitor))
}

func InfringementMonitorSetDealInfo(ctx echo.Context) error {
	idStr := ctx.QueryParam("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("获取id失败！", err.Error()))
	}
	dealInfo := &service.InfringementMonitorDealInfo{Id: id}
	if err := ctx.Bind(dealInfo); err != nil {
		return ctx.JSON(utils.ErrIpt("输入解析失败！", err.Error()))
	}
	if err := ctx.Validate(dealInfo); err != nil {
		return ctx.JSON(utils.ErrIpt("输入校验失败！", err.Error()))
	}
	if err := service.InfringementMonitorSetDealInfo(dealInfo); err != nil {
		return ctx.JSON(utils.ErrIpt("设置回访记录失败！", err.Error()))
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
	monitors, err := service.InfringementMonitorBackendList(page, search)
	if err != nil {
		return ctx.JSON(utils.ErrSvr("获取infringement_monitor list失败", err.Error()))
	}
	return ctx.JSON(utils.Succ("success", monitors))
}
