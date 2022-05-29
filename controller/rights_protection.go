package controller

import (
	dao "law/dao"
	"law/enum"
	"law/service"
	"law/utils"
	"strconv"

	"github.com/labstack/echo/v4"
)

func RightsProtectionAdd(ctx echo.Context) error {
	baseInfo := &service.CooperationBaseInfo{}
	if err := utils.BindAndValidate(ctx, baseInfo); err != nil {
		return ctx.JSON(utils.ErrIpt("输入解析校验失败！", err.Error()))
	}
	uid := ctx.Get("uid").(int)
	beanId, err := service.CooperationSrv.Add(baseInfo, enum.PROTECT, uid)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("添加维权意向失败！", err.Error()))
	}
	return ctx.JSON(utils.Succ("success", map[string]int{"rights_protection_id": beanId}))
}

func RightsProtectionBgGet(ctx echo.Context) error {
	beanIdStr := ctx.QueryParam("id")
	beanId, err := strconv.Atoi(beanIdStr)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("获取id失败", err.Error()))
	}
	bean, err := service.CooperationSrv.BgGet(beanId)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("获取侵权监测失败！", err.Error()))
	}
	return ctx.JSON(utils.Succ("success", bean))
}

func RightsProtectionGet(ctx echo.Context) error {
	uid := ctx.Get("uid").(int)
	bean, err := service.CooperationSrv.Get(enum.PROTECT, uid)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("查询失败！", err.Error()))
	}
	return ctx.JSON(utils.Succ("success", bean))
}

func RightsProtectionSetDealInfo(ctx echo.Context) error {
	dealInfo := &service.CooperationDealInfo{}
	beanIdStr := ctx.QueryParam("id")
	beanId, err := strconv.Atoi(beanIdStr)
	if err == nil {
		dealInfo.Id = beanId
	}
	if err := utils.BindAndValidate(ctx, dealInfo); err != nil {
		return ctx.JSON(utils.ErrIpt("输入解析校验失败！", err.Error()))
	}
	if err := service.CooperationSrv.SetDealInfo(dealInfo.Id, dealInfo); err != nil {
		return ctx.JSON(utils.ErrIpt("设置回访记录失败！", err.Error()))
	}
	return ctx.JSON(utils.Succ("success"))
}

func RightsProtectionUpdateBaseInfo(ctx echo.Context) error {
	baseInfo := &service.CooperationBaseInfo{}
	if err := utils.BindAndValidate(ctx, baseInfo); err != nil {
		return ctx.JSON(utils.ErrIpt("输入解析校验失败！", err.Error()))
	}
	uid := ctx.Get("uid").(int)
	if err := service.CooperationSrv.UpdateBaseInfo(enum.PROTECT, uid, baseInfo); err != nil {
		return ctx.JSON(utils.ErrIpt("修改基础信息失败！", err.Error()))
	}
	return ctx.JSON(utils.Succ("success"))
}

func RightsProtectionBackendList(ctx echo.Context) error {
	page := &dao.Page{PageIndex: 1, ItemNum: 10}
	if err := ctx.Bind(page); err != nil {
		return ctx.JSON(utils.ErrIpt("分页输入错误,请重试！", err.Error()))
	}
	if err := ctx.Validate(page); err != nil {
		return ctx.JSON(utils.ErrIpt("分页数据输入校验失败！", err.Error()))
	}
	search := &dao.CooperationSearchParams{}
	if err := ctx.Bind(search); err != nil {
		return ctx.JSON(utils.ErrIpt("检索数据输入错误,请重试！", err.Error()))
	}
	if err := ctx.Validate(search); err != nil {
		return ctx.JSON(utils.ErrIpt("检索输入校验失败！", err.Error()))
	}
	beans, err := dao.CooperationDao.BackendList(enum.PROTECT, page, search)
	if err != nil {
		return ctx.JSON(utils.ErrSvr("获取rights_protection list失败", err.Error()))
	}
	return ctx.JSON(utils.Succ("success", beans))
}
