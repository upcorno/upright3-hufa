package controller

import (
	"law/model"
	"law/service"
	"law/utils"
	"strconv"

	"github.com/labstack/echo/v4"
)

func LegalIssueList(ctx echo.Context) error {
	page := &model.Page{PageIndex: 1, ItemNum: 5}
	if err := ctx.Bind(page); err != nil {
		return ctx.JSON(utils.ErrIpt("分页输入错误,请重试！", err.Error()))
	}
	if err := ctx.Validate(page); err != nil {
		return ctx.JSON(utils.ErrIpt("分页数据输入校验失败！", err.Error()))
	}
	uid := ctx.Get("uid").(int)
	search := &model.LegalIssueSearch{FavoriteUid: uid, IsFavorite: false}
	if err := utils.BindAndValidate(ctx, search); err != nil {
		return ctx.JSON(utils.ErrIpt("检索数据输入错误,请重试！", err.Error()))
	}
	issues, err := model.LegalIssueList(page, search)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("获取legal_issue list失败！", err.Error()))
	}
	return ctx.JSON(utils.Succ("success", issues))
}

func LegalIssueGet(ctx echo.Context) error {
	legalIssueIdStr := ctx.QueryParam("legal_issue_id")
	legalIssueId, err := strconv.Atoi(legalIssueIdStr)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("获取legal_issue_id失败！", err.Error()))
	}
	issueInfo, err := service.LegalIssueSrv.GetLegalIssue(legalIssueId)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("获取LegalIssue失败！", err.Error()))
	}
	return ctx.JSON(utils.Succ("success", issueInfo))
}

func LegalIssueCategoryList(ctx echo.Context) error {
	categoryList, err := model.LegalIssueCategoryList()
	if err != nil {
		return ctx.JSON(utils.ErrIpt("获取问题分类列表失败！", err.Error()))
	}
	return ctx.JSON(utils.Succ("success", categoryList))
}

func LegalIssueFavorite(ctx echo.Context) error {
	issueIdStr := ctx.QueryParam("issue_id")
	issueId, err := strconv.Atoi(issueIdStr)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("获取issue_id失败！", err.Error()))
	}
	uid := ctx.Get("uid").(int)
	if err := service.LegalIssueSrv.AddFavorite(uid, issueId); err != nil {
		return ctx.JSON(utils.ErrIpt("添加收藏失败！", err.Error()))
	}
	return ctx.JSON(utils.Succ("success"))
}

func LegalIssueCancelFavorite(ctx echo.Context) error {
	issueIdStr := ctx.QueryParam("issue_id")
	issueId, err := strconv.Atoi(issueIdStr)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("获取issue_id失败！", err.Error()))
	}
	uid := ctx.Get("uid").(int)
	if err := service.LegalIssueSrv.CancelFavorite(uid, issueId); err != nil {
		return ctx.JSON(utils.ErrIpt("取消收藏失败！", err.Error()))
	}
	return ctx.JSON(utils.Succ("success"))
}

func LegalIssueIsFavorite(ctx echo.Context) error {
	issueIdStr := ctx.QueryParam("issue_id")
	issueId, err := strconv.Atoi(issueIdStr)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("获取issue_id失败！", err.Error()))
	}
	uid := ctx.Get("uid").(int)
	has, err := service.LegalIssueSrv.IsFavorite(uid, issueId)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("查询普法问题收藏失败！", err.Error()))
	}
	return ctx.JSON(utils.Succ("success", has))
}
