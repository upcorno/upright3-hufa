package service

import (
	"encoding/json"
	dao "law/dao"
	"law/enum"

	zlog "github.com/rs/zerolog/log"
)

type consultationSrv struct {
}

var Consultation *consultationSrv = &consultationSrv{}

type ConsultationCreateInfo struct {
	Question string `json:"question" validate:"required"`
	Imgs     string `json:"imgs"`
}

func (c *consultationSrv) Create(createInfo *ConsultationCreateInfo, consultantUid int) (consulId int, err error) {
	if consulId, err = dao.ConsulDao.Insert(createInfo.Question, createInfo.Imgs, consultantUid, enum.DOING); err != nil {
		return
	}
	consultationData, err := json.Marshal(createInfo)
	if err != nil {
		return
	}
	replyParams := &ConsultationReplyParams{
		ConsultationId: consulId,
		Type:           enum.QUERY,
		Content:        string(consultationData),
	}
	err = c.AddReply(replyParams, consultantUid, enum.NO)
	return
}

func (c *consultationSrv) SetStatus(consulId int, status string) (err error) {
	consul := &dao.Consultation{Status: status}
	err = dao.ConsulDao.Update(consulId, consul, "status")
	return err
}

type ConsultationReplyParams struct {
	ConsultationId int    `json:"consultationId" form:"consultationId" query:"consultationId" validate:"required,min=1"`
	Type           string `json:"type" validate:"required,oneof=answer query"`
	Content        string `json:"content" validate:"required,min=1"`
}

func (c *consultationSrv) AddReply(replyParams *ConsultationReplyParams, uid int, isMannerReply enum.YesOrNo) (err error) {
	_, err = dao.ConsulReplyDao.Insert(
		replyParams.ConsultationId,
		replyParams.Type,
		replyParams.Content,
		uid,
		isMannerReply,
	)
	go func() {
		if isMannerReply == enum.YES {
			mannerReplys, err := dao.ConsulReplyDao.List(replyParams.ConsultationId, 0, true)
			if err != nil {
				zlog.Error().Msg("查询人工回复记录失败。" + err.Error())
				return
			}
			consul, err := dao.ConsulDao.GetWithUserInfo(replyParams.ConsultationId)
			if err != nil {
				zlog.Error().Msg("查询Consultation失败。" + err.Error())
				return
			}
			if len(mannerReplys) == 1 {
				WxSrv.SendConsulNotify(consul.ConsultantUid, replyParams.ConsultationId, consul.Question, consul.CreateTime)
			}
		}
	}()
	return
}
