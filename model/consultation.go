package model

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type ConsultationInfo struct {
	Id            int       `json:"id"`
	Question      string    `json:"question"`
	Imgs          string    `json:"imgs"`
	ConsultantUid int       `json:"consultant_uid"`
	NickName      string    `json:"nick_name"`
	AvatarUrl     string    `json:"avatar_url"`
	Phone         string    `json:"phone"`
	Status        string    `json:"status"`
	CreateTime    int       `json:"create_time"`
	UpdateTime    time.Time `json:"update_time"`
}

//创建咨询
func ConsultationCreate(consul *Consultation) error {
	_, err := Db.InsertOne(consul)
	return err
}

//设置咨询状态
func ConsultationStatusSet(consultationId int, status string) error {
	_, err := Db.Cols("status").Update(&Consultation{Status: status}, &Consultation{Id: consultationId})
	return err
}

//用户历史咨询记录列表
func ConsultationList(uid int) ([]Consultation, error) {
	consultationList := []Consultation{}
	err := Db.Table("consultation").
		Where("consultation.consultant_uid = ?", uid).
		Find(&consultationList)
	return consultationList, err
}

//获取咨询信息
func ConsultationInfoGet(consultationId int) (ConsultationInfo, error) {
	consultationInfo := ConsultationInfo{}
	_, err := Db.Table("consultation").
		Join("INNER", "user", "user.id = consultation.consultant_uid").
		Where("consultation.id=?", consultationId).
		Cols(
			"consultation.id",
			"consultation.question",
			"consultation.imgs",
			"consultation.consultant_uid",
			"user.nick_name",
			"user.avatar_url",
			"user.phone",
			"consultation.status",
			"consultation.create_time",
			"consultation.update_time",
		).
		Get(&consultationInfo)
	return consultationInfo, err
}
