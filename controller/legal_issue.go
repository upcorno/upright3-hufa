package controller

import (
	"law/model"
	"law/utils"
	"strconv"

	"github.com/labstack/echo/v4"
)

//普法问题列表检索
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
	if err := ctx.Bind(search); err != nil {
		return ctx.JSON(utils.ErrIpt("检索数据输入错误,请重试！", err.Error()))
	}
	if err := ctx.Validate(search); err != nil {
		return ctx.JSON(utils.ErrIpt("检索输入校验失败！", err.Error()))
	}
	issues, err := model.LegalIssueList(page, search)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("获取legal_issue list失败！", err.Error()))
	}
	return ctx.JSON(utils.Succ("success", issues))
}

//获取普法问题
func LegalIssueGet(ctx echo.Context) error {
	legalIssueIdStr := ctx.QueryParam("legal_issue_id")
	legalIssueId, err := strconv.Atoi(legalIssueIdStr)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("获取legal_issue_id失败！", err.Error()))
	}
	issue, err := model.LegalIssueGet(legalIssueId)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("获取普法问题失败！", err.Error()))
	}
	return ctx.JSON(utils.Succ("success", issue))
}

func LegalIssueCategoryList(ctx echo.Context) error {
	categoryList, err := model.IssueCategoryList()
	if err != nil {
		return ctx.JSON(utils.ErrIpt("获取问题分类列表失败！", err.Error()))
	}
	return ctx.JSON(utils.Succ("success", categoryList))
}

//收藏普法问题
func LegalIssueFavorite(ctx echo.Context) error {
	issueIdStr := ctx.QueryParam("issue_id")
	issueId, err := strconv.Atoi(issueIdStr)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("获取issue_id失败！", err.Error()))
	}
	uid := ctx.Get("uid").(int)
	if err := model.LegalIssueAddFavorite(uid, issueId); err != nil {
		return ctx.JSON(utils.ErrIpt("添加收藏失败！", err.Error()))
	}
	return ctx.JSON(utils.Succ("success"))
}

//取消收藏普法问题
func LegalIssueCancelFavorite(ctx echo.Context) error {
	issueIdStr := ctx.QueryParam("issue_id")
	issueId, err := strconv.Atoi(issueIdStr)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("获取issue_id失败！", err.Error()))
	}
	uid := ctx.Get("uid").(int)
	if err := model.LegalIssueCancelFavorite(uid, issueId); err != nil {
		return ctx.JSON(utils.ErrIpt("取消收藏失败！", err.Error()))
	}
	return ctx.JSON(utils.Succ("success"))
}

//普法问题是否收藏
func LegalIssueIsFavorite(ctx echo.Context) error {
	issueIdStr := ctx.QueryParam("issue_id")
	issueId, err := strconv.Atoi(issueIdStr)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("获取issue_id失败！", err.Error()))
	}
	uid := ctx.Get("uid").(int)
	has, err := model.LegalIssueIsFavorite(uid, issueId)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("查询普法问题收藏失败！", err.Error()))
	}
	return ctx.JSON(utils.Succ("success", has))
}
