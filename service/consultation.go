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

var Consultation *consultationSrv

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
	consul.CreateTime = int(time.Now().Unix())
	if err = model.ConsultationCreate(consul); err != nil {
		return
	}
	consultationId = consul.Id
	consultationData, err := json.Marshal(createInfo)
	if err != nil {
		return
	}
	record := &model.ConsultationReply{
		CommunicatorUid: uid,
		Type:            enum.QUERY,
		Content:         string(consultationData),
		CreateTime:      int(time.Now().Unix()),
	}
	err = consul.AddReply(record)
	return
}

type ConsultationSearchParams struct {
	Status        string `json:"status" query:"status"`
	CreateTimeMin int    `json:"create_time_min" query:"create_time_min"`
	CreateTimeMax int    `json:"create_time_max" query:"create_time_max"`
}

//侵权监测列表搜索
func (c *consultationSrv) BackendList(page *model.Page, search *ConsultationSearchParams) (*model.PageResult, error) {
	type listInfo struct {
		Id       int    `json:"id"`
		Question string `json:"question"`
		NickName string `json:"nick_name"`
		Phone    string `json:"phone"`
		Status   string `json:"status"`
	}
	consultationInfo := []listInfo{}
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
	pageResult, err := page.GetResults(sess, &consultationInfo)
	if err != nil {
		return nil, err
	}
	return pageResult, err
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
