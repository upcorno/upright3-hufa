package controller

import (
	"law/model"
	"law/service"
	"law/utils"

	"github.com/labstack/echo/v4"
)

func LegalIssueList(ctx echo.Context) error {
	page := &model.Page{PageIndex:1, ItemNum: 5}
	if err := ctx.Bind(page); err != nil {
		return ctx.JSON(utils.ErrIpt("分页输入错误,请重试！", err.Error()))
	}
	if err := ctx.Validate(page); err != nil {
		return ctx.JSON(utils.ErrIpt("分页数据输入校验失败！", err.Error()))
	}
	search := &service.LegalIssueSearch{}
	if err := ctx.Bind(search); err != nil {
		return ctx.JSON(utils.ErrIpt("检索数据输入错误,请重试！", err.Error()))
	}
	if err := ctx.Validate(search); err != nil {
		return ctx.JSON(utils.ErrIpt("检索输入校验失败！", err.Error()))
	}
	issues, err := service.LegalIssueList(page, search)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("获取legal_issue list失败！", err.Error()))
	}
	return ctx.JSON(utils.Succ("success",issues))
}