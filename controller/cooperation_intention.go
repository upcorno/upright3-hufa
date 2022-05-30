package controller

import (
	"errors"
	dao "law/dao"
	"law/enum"
	"law/service"
	"law/utils"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type cooperationController struct{}

var CooperationController *cooperationController

func (c *cooperationController) getCooperationType(path string) (category enum.Cooperation, err error) {
	if strings.Contains(path, "/infringement_monitor/") {
		category = enum.MONITOR
		return
	} else if strings.Contains(path, "/rights_protection/") {
		category = enum.PROTECT
		return
	}
	err = errors.New("不支持的合作类型")
	return
}

func (c *cooperationController) Add(ctx echo.Context) error {
	baseInfo := &service.CooperationBaseInfo{}
	if err := utils.BindAndValidate(ctx, baseInfo); err != nil {
		return ctx.JSON(utils.ErrIpt("输入解析校验失败！", err.Error()))
	}
	uid := ctx.Get("uid").(int)
	category, err := c.getCooperationType(ctx.Path())
	if err != nil {
		return ctx.JSON(utils.ErrIpt(err.Error()))
	}
	beanId, err := service.CooperationSrv.Add(baseInfo, category, uid)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("添加合作意向失败！", err.Error()))
	}
	data := map[string]int{"id": beanId}
	///兼容旧版，以后可删除
	if category == enum.MONITOR {
		data["infringement_monitor_id"] = beanId
	}
	if category == enum.PROTECT {
		data["rights_protection_id"] = beanId
	}
	///
	return ctx.JSON(utils.Succ("success", data))
}

func (c *cooperationController) Get(ctx echo.Context) error {
	uid := ctx.Get("uid").(int)
	category, err := c.getCooperationType(ctx.Path())
	if err != nil {
		return ctx.JSON(utils.ErrIpt(err.Error()))
	}
	bean, err := service.CooperationSrv.Get(category, uid)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("获取合作意向失败！", err.Error()))
	}
	return ctx.JSON(utils.Succ("success", bean))
}

func (c *cooperationController) BgGet(ctx echo.Context) error {
	beanIdStr := ctx.QueryParam("id")
	beanId, err := strconv.Atoi(beanIdStr)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("获取id失败", err.Error()))
	}
	category, err := c.getCooperationType(ctx.Path())
	if err != nil {
		return ctx.JSON(utils.ErrIpt(err.Error()))
	}
	bean, err := service.CooperationSrv.BgGet(category, beanId)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("获取合作意向失败！", err.Error()))
	}
	return ctx.JSON(utils.Succ("success", bean))
}

func (c *cooperationController) SetDealInfo(ctx echo.Context) error {
	dealInfo := &service.CooperationDealInfo{}
	beanIdStr := ctx.QueryParam("id")
	beanId, err := strconv.Atoi(beanIdStr)
	if err == nil {
		dealInfo.Id = beanId
	}
	if err := utils.BindAndValidate(ctx, dealInfo); err != nil {
		return ctx.JSON(utils.ErrIpt("输入解析校验失败！", err.Error()))
	}
	category, err := c.getCooperationType(ctx.Path())
	if err != nil {
		return ctx.JSON(utils.ErrIpt(err.Error()))
	}
	if err := service.CooperationSrv.SetDealInfo(category, dealInfo.Id, dealInfo); err != nil {
		return ctx.JSON(utils.ErrIpt("设置回访记录失败！", err.Error()))
	}
	return ctx.JSON(utils.Succ("success"))
}

func (c *cooperationController) UpdateBaseInfo(ctx echo.Context) error {
	baseInfo := &service.CooperationBaseInfo{}
	if err := utils.BindAndValidate(ctx, baseInfo); err != nil {
		return ctx.JSON(utils.ErrIpt("输入解析校验失败！", err.Error()))
	}
	uid := ctx.Get("uid").(int)
	category, err := c.getCooperationType(ctx.Path())
	if err != nil {
		return ctx.JSON(utils.ErrIpt(err.Error()))
	}
	if err := service.CooperationSrv.UpdateBaseInfo(category, uid, baseInfo); err != nil {
		return ctx.JSON(utils.ErrIpt("修改基础信息失败！", err.Error()))
	}
	return ctx.JSON(utils.Succ("success"))
}

func (c *cooperationController) BackendList(ctx echo.Context) error {
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
	category, err := c.getCooperationType(ctx.Path())
	if err != nil {
		return ctx.JSON(utils.ErrIpt(err.Error()))
	}
	beans, err := dao.CooperationDao.BackendList(category, page, search)
	if err != nil {
		return ctx.JSON(utils.ErrSvr("获取infringement_monitor list失败", err.Error()))
	}
	return ctx.JSON(utils.Succ("success", beans))
}
