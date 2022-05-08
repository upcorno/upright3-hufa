package model

import (
	"errors"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// 问题“咨询”
type Consultation struct {
	Id            int       `xorm:"not null pk autoincr UNSIGNED INT" json:"id"`
	Question      string    `xorm:"not null comment('咨询问题') TEXT" json:"question" validate:"required"`
	Imgs          string    `xorm:"comment('描述图片') TEXT" json:"imgs"`
	ConsultantUid int       `xorm:"not null comment('咨询人uid') index UNSIGNED INT" json:"consultant_uid" validate:"required"`
	Status        string    `xorm:"not null default '处理中' comment('处理中、待人工咨询、人工咨询中、已完成') VARCHAR(10)" json:"status" validate:"required,oneof=处理中 待人工咨询 人工咨询中 已完成"`
	CreateTime    int       `xorm:"not null UNSIGNED INT" json:"create_time"`
	UpdateTime    time.Time `xorm:"not null updated DateTime default(CURRENT_TIMESTAMP)" json:"-"`
}

//创建咨询
func ConsultationCreate(consul *Consultation) error {
	id, err := Db.InsertOne(consul)
	consul.Id = int(id)
	return err
}

//设置咨询状态
func (consul *Consultation) SetStatus(status string) error {
	if consul.Id == 0 {
		err := errors.New("model:必须指定id值")
		return err
	}
	_, err := Db.Cols("status").Update(
		Consultation{Status: status},
		Consultation{Id: consul.Id},
	)
	return err
}

//创建咨询记录
func (consul *Consultation) AddReply(record *ConsultationReply) error {
	if consul.Id == 0 {
		err := errors.New("model:必须指定id值")
		return err
	}
	record.ConsultationId = consul.Id
	_, err := Db.InsertOne(record)
	return err
}

//获取咨询沟通记录表
func (consul *Consultation) ListReply(consultationId int) (recordInfoList []map[string]string, err error) {
	if consul.Id == 0 {
		err = errors.New("model:必须指定id值")
		return
	}
	err = Db.Table("consultation_reply").
		Join("INNER", "user", "user.id = consultation_reply.communicator_uid").
		Where("consultation_id=?", consultationId).
		Cols(
			"consultation_reply.id",
			"consultation_id",
			"communicator_uid",
			"type",
			"content",
			"nick_name",
			"avatar_url",
			"phone",
			"consultation_reply.create_time",
		).
		Asc("consultation_reply.create_time").
		Find(&recordInfoList)
	return
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
func ConsultationGetWithUserInfo(consultationId int) (map[string]string, error) {
	consultationInfo := map[string]string{}
	_, err := Db.Table("consultation").
		Join("INNER", "user", "user.id = consultation.consultant_uid").
		Where("consultation.id=?", consultationId).
		Cols(
			"consultation.id",
			"consultation.question",
			"consultation.imgs",
			"consultation.status",
			"consultation.create_time",
			"consultation.update_time",
			"consultation.consultant_uid",
			"user.nick_name",
			"user.avatar_url",
			"user.phone",
		).
		Get(&consultationInfo)
	if _, ok := consultationInfo["id"]; !ok {
		return nil, err
	}
	return consultationInfo, err
}
