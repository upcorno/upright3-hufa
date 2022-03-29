package controller

import (
	"encoding/json"
	"law/enum"
	"law/model"
	"law/utils"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

//创建咨询实例
func ConsultationCreate(ctx echo.Context) error {
	consul := &model.Consultation{}
	if err := ctx.Bind(consul); err != nil {
		return ctx.JSON(utils.ErrIpt("输入解析失败！", err.Error()))
	}
	uid := ctx.Get("uid").(int)
	consul.ConsultantUid = uid
	consul.Status = enum.DOING
	consul.CreateTime = int(time.Now().Unix())
	if err := ctx.Validate(consul); err != nil {
		return ctx.JSON(utils.ErrIpt("输入校验失败！", err.Error()))
	}
	if err := model.ConsultationCreate(consul); err != nil {
		return ctx.JSON(utils.ErrIpt("法律咨询生成失败！", err.Error()))
	}
	consultationData, err := json.Marshal(map[string]string{"question":consul.Question, "imgs":consul.Imgs})
	if err != nil {
		return ctx.JSON(utils.ErrOpt("consultation info序列化失败", err.Error()))
	}
	record := &model.ConsultationRecord{
		ConsultationId: consul.Id,
		CommunicatorUid: uid,
		Type: enum.QUERY,
		Content: string(consultationData),
		CreateTime: int(time.Now().Unix()),
	}
	if err := model.ConsultationRecordCreate(record); err != nil {
		return ctx.JSON(utils.ErrIpt("法律咨询记录生成失败！", err.Error()))
	}
	return ctx.JSON(utils.Succ("success", map[string]int{"consultation_id": consul.Id}))
}

//咨询设置状态
func ConsultationStatusSet(ctx echo.Context) error {
	consultationIdStr := ctx.QueryParam("consultation_id")
	consultationId, err := strconv.Atoi(consultationIdStr)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("获取consultation_id失败！", err.Error()))
	}
	status := ctx.QueryParam("status")
	if err := model.ConsultationStatusSet(consultationId, status); err != nil {
		return ctx.JSON(utils.ErrIpt("法律咨询状态设置失败！", err.Error()))
	}
	return ctx.JSON(utils.Succ("success", map[string]int{"consultation_id": consultationId}))
}

//用户历史咨询记录
func ConsultationList(ctx echo.Context) error {
	uid := ctx.Get("uid").(int)
	consultationList, err := model.ConsultationList(uid)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("获取历史咨询列表失败！", err.Error()))
	}
	return ctx.JSON(utils.Succ("success", consultationList))
}

//咨询文件上传
func ConsultationFileUploadAuth(ctx echo.Context) error {
	fileUploadConfInfo, err := utils.FileUploadAuthSTS("consultation")
	if err != nil {
		return ctx.JSON(utils.ErrIpt("获取文件上传配置信息失败", err.Error()))
	}
	return ctx.JSON(utils.Succ("success", fileUploadConfInfo))
}