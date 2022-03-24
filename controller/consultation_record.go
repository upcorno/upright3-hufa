package controller

import (
	"law/model"
	"law/utils"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

//创建咨询记录
func ConsultationRecordCreate(ctx echo.Context) error {
	consultationIdStr := ctx.QueryParam("consultation_id")
	consultationId, err := strconv.Atoi(consultationIdStr)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("获取consultation_id失败！", err.Error()))
	}
	record := &model.ConsultationRecord{}
	if err := ctx.Bind(record); err != nil {
		return ctx.JSON(utils.ErrIpt("输入解析失败！", err.Error()))
	}
	record.ConsultationId = consultationId
	uid := ctx.Get("uid").(int)
	record.CommunicatorUid = uid
	record.CreateTime = int(time.Now().Unix())
	if err := ctx.Validate(record); err != nil {
		return ctx.JSON(utils.ErrIpt("输入校验失败！", err.Error()))
	}
	if err := model.ConsultationRecordCreate(record); err != nil {
		return ctx.JSON(utils.ErrIpt("法律咨询记录生成失败！", err.Error()))
	}
	return ctx.JSON(utils.Succ("success"))
}

//获取咨询沟通记录
func ConsultationRecordListGet(ctx echo.Context) error {
	consultationIdStr := ctx.QueryParam("consultation_id")
	consultationId, err := strconv.Atoi(consultationIdStr)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("获取consultation_id失败！", err.Error()))
	}
	recordList, err := model.ConsultationRecordList(consultationId)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("获取consultation_record_list失败！", err.Error()))
	}
	return ctx.JSON(utils.Succ("success", recordList))
}