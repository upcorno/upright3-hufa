package controller

import (
	"law/enum"
	"law/model"
	"law/utils"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

//添加侵权监测
func InfringementDetectionAdd(ctx echo.Context) error {
	infringementDetection := &model.InfringementDetection{}
	if err := ctx.Bind(infringementDetection); err != nil {
		return ctx.JSON(utils.ErrIpt("输入解析失败！", err.Error()))
	}
	uid := ctx.Get("uid").(int)
	infringementDetection.CreatorUid = uid
	infringementDetection.CreateTime = int(time.Now().Unix())
	err := model.InfringementDetectionAdd(infringementDetection)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("添加侵权监测失败！", err.Error()))
	}
	err = model.DetectionRetureVisitAdd(&model.DetectionReturnVisit{
		CreatorUid: uid,
		DetectionId: infringementDetection.Id,
		Classification: enum.NORETURN,
		CreateTime: int(time.Now().Unix()),
	})
	if err != nil {
		return ctx.JSON(utils.ErrIpt("添加默认回访失败！", err.Error()))
	}
	return ctx.JSON(utils.Succ("success", map[string]int{"infringement_detection_id": infringementDetection.Id}))
}

//获取侵权监测
func InfringementDetectionGet(ctx echo.Context) error {
	detectionIdStr := ctx.QueryParam("detection_id")
	detectionId, err := strconv.Atoi(detectionIdStr)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("获取detection_id失败！", err.Error()))
	}
	infringementDetection, err := model.InfringementDetectionGet(detectionId)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("获取侵权监测失败！", err.Error()))
	}
	return ctx.JSON(utils.Succ("success", infringementDetection))
}