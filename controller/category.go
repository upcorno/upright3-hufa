package controller

import (
	"law/model"
	"law/utils"
	"strconv"

	"github.com/labstack/echo/v4"
)


//获取当前分类子级分类
func CategoryListGet(ctx echo.Context) error {
	categoryIdStr := ctx.QueryParam("category_id")
	categoryId, err := strconv.Atoi(categoryIdStr)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("获取category_id失败！", err.Error()))
	}
	categoryList, err := model.CategoryList(categoryId)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("获取consultation_record_list失败！", err.Error()))
	}
	return ctx.JSON(utils.Succ("success", categoryList))
}