package service

import (
	"encoding/json"
	dao "law/dao"
	"law/enum"
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
	err = c.AddReply(replyParams, consultantUid)
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

func (c *consultationSrv) AddReply(replyParams *ConsultationReplyParams, uid int) (err error) {
	_, err = dao.ConsulReplyDao.Insert(
		replyParams.ConsultationId,
		replyParams.Type,
		replyParams.Content,
		uid,
	)
	return
}
