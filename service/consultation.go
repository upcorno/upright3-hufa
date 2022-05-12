package service

import (
	"encoding/json"
	"law/enum"
	"law/model"
	"time"

	"xorm.io/xorm"
)

type consultationSrv struct {
}

var Consultation *consultationSrv = &consultationSrv{}

type ConsultationCreateInfo struct {
	Question string `json:"question" validate:"required"`
	Imgs     string `json:"imgs"`
}

func (c *consultationSrv) Create(createInfo *ConsultationCreateInfo, uid int) (consultationId int, err error) {
	consul := &model.Consultation{
		Question: createInfo.Question,
		Imgs:     createInfo.Imgs,
	}
	consul.ConsultantUid = uid
	consul.Status = enum.DOING
	if err = consul.Create(); err != nil {
		return
	}
	consultationId = consul.Id
	consultationData, err := json.Marshal(createInfo)
	if err != nil {
		return
	}
	reply := &model.ConsultationReply{
		ConsultationId:  consultationId,
		Type:            enum.QUERY,
		Content:         string(consultationData),
		CommunicatorUid: uid,
		CreateTime:      int(time.Now().Unix()),
	}
	err = reply.Insert()
	return
}
func (c *consultationSrv) SetStatus(consultationId int, status string) (err error) {
	consul := &model.Consultation{Id: consultationId, Status: status}
	err = consul.Update("status")
	return err
}

type ConsultationReplyParams struct {
	ConsultationId int    `json:"consultationId" form:"consultationId" query:"consultationId" validate:"required,min=1"`
	Type           string `json:"type" validate:"required,oneof=answer query"`
	Content        string `json:"content" validate:"required,min=1"`
}

func (c *consultationSrv) AddReply(replyParams *ConsultationReplyParams, uid int) (err error) {
	reply := &model.ConsultationReply{
		ConsultationId:  replyParams.ConsultationId,
		Type:            replyParams.Type,
		Content:         replyParams.Content,
		CommunicatorUid: uid,
		CreateTime:      int(time.Now().Unix()),
	}
	err = reply.Insert()
	return
}

type ConsultationSearchParams struct {
	Status        string `json:"status" query:"status"`
	CreateTimeMin int    `json:"create_time_min" query:"create_time_min"`
	CreateTimeMax int    `json:"create_time_max" query:"create_time_max"`
}

func (c *consultationSrv) BackendList(page *model.Page, search *ConsultationSearchParams) (pageResult *model.PageResult, err error) {
	type listInfo struct {
		Id       int    `json:"id"`
		Question string `json:"question"`
		NickName string `json:"nick_name"`
		Phone    string `json:"phone"`
		Status   string `json:"status"`
	}
	consultationInfoList := []listInfo{}
	sess := model.Db.NewSession()
	sess.Table("consultation")
	sess.Join("INNER", "user", "user.id = consultation.consultant_uid")
	sess.Cols(
		"consultation.id",
		"consultation.question",
		"user.nick_name",
		"user.phone",
		"consultation.status",
	)
	c.dealSearch(sess, search)
	pageResult, err = page.GetResults(sess, &consultationInfoList)
	return
}

func (c *consultationSrv) dealSearch(sess *xorm.Session, search *ConsultationSearchParams) {
	if search.Status != "" {
		sess.Where("consultation.status = ?", search.Status)
	}
	if search.CreateTimeMin != 0 {
		sess.Where("consultation.create_time > ?", search.CreateTimeMin)
	}
	if search.CreateTimeMax != 0 {
		sess.Where("consultation.create_time < ?", search.CreateTimeMax)
	}
}
