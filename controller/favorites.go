package controller

import (
	"law/model"
	"law/utils"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

//收藏普法问题
func FavoritesAdd(ctx echo.Context) error {
	issueIdStr := ctx.QueryParam("issue_id")
	issueId, err := strconv.Atoi(issueIdStr)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("获取issue_id失败！", err.Error()))
	}
	favorites := &model.Favorites{}
	if err := ctx.Bind(favorites); err != nil {
		return ctx.JSON(utils.ErrIpt("输入解析失败！", err.Error()))
	}
	favorites.IssueId = issueId
	uid := ctx.Get("uid").(int)
	favorites.UserId = uid
	favorites.CreateTime = int(time.Now().Unix())
	if err := ctx.Validate(favorites); err != nil {
		return ctx.JSON(utils.ErrIpt("输入校验失败！", err.Error()))
	}
	if err := model.FavoritesAdd(favorites); err != nil {
		return ctx.JSON(utils.ErrIpt("添加收藏失败！", err.Error()))
	}
	return ctx.JSON(utils.Succ("success"))
}

//取消收藏普法问题
func FavoritesCancel(ctx echo.Context) error {
	issueIdStr := ctx.QueryParam("issue_id")
	issueId, err := strconv.Atoi(issueIdStr)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("获取issue_id失败！", err.Error()))
	}
	uid := ctx.Get("uid").(int)
	if err := model.FavoritesCancel(uid, issueId); err != nil {
		return ctx.JSON(utils.ErrIpt("取消收藏失败！", err.Error()))
	}
	return ctx.JSON(utils.Succ("success"))
}