package service

import (
	"law/model"

	"xorm.io/xorm"
)

type ConsultationSearch struct {
	Status  string `json:"status" query:"status"`
	CreateTimeMin   int    `json:"create_time_min" query:"create_time_min"`
	CreateTimeMax   int    `json:"create_time_max" query:"create_time_max"`
}

type ConsultationInfo struct {
	Id       int    `json:"id"`
	Question string `json:"question"`   
	NickName string `json:"nick_name"`
	Phone    string `json:"phone"`
	Status   string `json:"status"`
}

//侵权监测列表搜索
func ConsultationSearchList(page *model.Page, search *ConsultationSearch) (*model.PageResult, error) {
	consultationInfo := []ConsultationInfo{}
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
	dealConsultationSearch(sess, search)
	pageResult, err := page.GetResults(sess, &consultationInfo)
	if err != nil {
		return nil, err
	}
	return pageResult, err
}

func dealConsultationSearch(sess *xorm.Session, search *ConsultationSearch) {
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
