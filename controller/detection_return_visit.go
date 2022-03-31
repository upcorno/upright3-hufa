package controller

import (
	"law/model"
	"law/utils"
	"strconv"

	"github.com/labstack/echo/v4"
)

//编辑回访记录
func DetectionReturnVisitUpdate(ctx echo.Context) error {
	detectionIdStr := ctx.QueryParam("detection_id")
	detectionId, err := strconv.Atoi(detectionIdStr)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("获取detection_id失败！", err.Error()))
	}
	returnVisit := &model.DetectionReturnVisit{}
	if err := ctx.Bind(returnVisit); err != nil {
		return ctx.JSON(utils.ErrIpt("输入解析失败！", err.Error()))
	}
	returnVisit.DetectionId = detectionId
	if err := ctx.Validate(returnVisit); err != nil {
		return ctx.JSON(utils.ErrIpt("输入校验失败！", err.Error()))
	}
	if err := model.DetectionRetureVisitUpdate(returnVisit); err != nil {
		return ctx.JSON(utils.ErrIpt("更新回访记录失败！", err.Error()))
	}
	return ctx.JSON(utils.Succ("success"))
}

//获取回访记录
func DetectionReturnVisitGet(ctx echo.Context) error {
	detectionIdStr := ctx.QueryParam("detection_id")
	detectionId, err := strconv.Atoi(detectionIdStr)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("获取detection_id失败！", err.Error()))
	}
	returnVisit, err := model.DetectionReturnVisitGet(detectionId)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("获取侵权监测回访记录失败！", err.Error()))
	}
	return ctx.JSON(utils.Succ("success", returnVisit))
}