package controller

import (
	"fmt"
	"law/model"
	"law/service"
	"law/utils"
	"strconv"

	"github.com/labstack/echo/v4"
)

//创建咨询实例
func ConsultationCreate(ctx echo.Context) error {
	consulCreateInfo := &service.ConsultationCreateInfo{}
	if err := utils.BindAndValidate(ctx, consulCreateInfo); err != nil {
		return ctx.JSON(utils.ErrIpt("输入解析失败！", err.Error()))
	}
	uid := ctx.Get("uid").(int)
	consulId, err := service.Consultation.Create(consulCreateInfo, uid)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("插入咨询记录失败", err.Error()))
	}
	return ctx.JSON(utils.Succ("success", map[string]int{"consultation_id": consulId}))
}

//咨询设置状态
func ConsultationSetStatus(ctx echo.Context) error {
	consultationIdStr := ctx.QueryParam("consultation_id")
	consultationId, err := strconv.Atoi(consultationIdStr)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("获取consultation_id失败！", err.Error()))
	}
	consul := &model.Consultation{Id: consultationId}
	status := ctx.QueryParam("status")
	if err := consul.SetStatus(status); err != nil {
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

//获取咨询信息
func ConsultationGet(ctx echo.Context) error {
	consultationIdStr := ctx.QueryParam("consultation_id")
	consultationId, err := strconv.Atoi(consultationIdStr)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("获取consultation_id失败！", err.Error()))
	}
	consultationInfo, err := model.ConsultationGetWithUserInfo(consultationId)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("获取咨询信息失败！", err.Error()))
	}
	if consultationInfo == nil {
		return ctx.JSON(utils.ErrIpt(fmt.Sprintf("不存在此咨询。ID：%d", consultationId)))
	}
	return ctx.JSON(utils.Succ("success", consultationInfo))
}

//咨询后台列表检索
func ConsultationBackendList(ctx echo.Context) error {
	page := &model.Page{PageIndex: 1, ItemNum: 10}
	if err := ctx.Bind(page); err != nil {
		return ctx.JSON(utils.ErrIpt("分页输入错误,请重试！", err.Error()))
	}
	if err := ctx.Validate(page); err != nil {
		return ctx.JSON(utils.ErrIpt("分页数据输入校验失败！", err.Error()))
	}
	search := &service.ConsultationSearchParams{}
	if err := ctx.Bind(search); err != nil {
		return ctx.JSON(utils.ErrIpt("检索数据输入错误,请重试！", err.Error()))
	}
	if err := ctx.Validate(search); err != nil {
		return ctx.JSON(utils.ErrIpt("检索输入校验失败！", err.Error()))
	}
	consultations, err := service.Consultation.BackendList(page, search)
	if err != nil {
		return ctx.JSON(utils.ErrSvr("获取consultation list失败", err.Error()))
	}
	return ctx.JSON(utils.Succ("success", consultations))
}

//创建咨询回复记录
func ConsultationAddReply(ctx echo.Context) error {
	consultationIdStr := ctx.QueryParam("consultation_id")
	consultationId, err := strconv.Atoi(consultationIdStr)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("获取consultation_id失败！", err.Error()))
	}
	replyParams := &service.ConsultationReplyParams{
		ConsultationId: consultationId,
	}
	if err := utils.BindAndValidate(ctx, replyParams); err != nil {
		return ctx.JSON(utils.ErrIpt("输入解析失败！", err.Error()))
	}
	uid := ctx.Get("uid").(int)
	if err := service.Consultation.AddReply(replyParams, uid); err != nil {
		return ctx.JSON(utils.ErrIpt("法律咨询回复添加失败！", err.Error()))
	}
	return ctx.JSON(utils.Succ("success"))
}

//获取咨询回复记录
func ConsultationListReply(ctx echo.Context) error {
	consultationIdStr := ctx.QueryParam("consultation_id")
	consultationId, err := strconv.Atoi(consultationIdStr)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("获取consultation_id失败！", err.Error()))
	}
	recordListInfo, err := model.ConsultationReplyList(consultationId)
	if err != nil {
		return ctx.JSON(utils.ErrIpt("获取consultation_reply_info list失败！", err.Error()))
	}
	return ctx.JSON(utils.Succ("success", recordListInfo))
}

//后台咨询回复文件上传
func ConsultationBackendFileUploadAuth(ctx echo.Context) error {
	fileUploadConfInfo, err := utils.FileUploadAuthSTS("consultation_backend")
	if err != nil {
		return ctx.JSON(utils.ErrIpt("获取文件上传配置信息失败", err.Error()))
	}
	return ctx.JSON(utils.Succ("success", fileUploadConfInfo))
}

//用户咨询文件上传
func ConsultationFileUploadAuth(ctx echo.Context) error {
	fileUploadConfInfo, err := utils.FileUploadAuthSTS("consultation")
	if err != nil {
		return ctx.JSON(utils.ErrIpt("获取文件上传配置信息失败", err.Error()))
	}
	return ctx.JSON(utils.Succ("success", fileUploadConfInfo))
}
