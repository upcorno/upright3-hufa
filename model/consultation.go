package model

import (
	"errors"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// 问题“咨询”
type Consultation struct {
	Id            int       `xorm:"not null pk autoincr UNSIGNED INT" json:"id"`
	Question      string    `xorm:"not null comment('咨询问题') TEXT" json:"question"`
	Imgs          string    `xorm:"not null comment('描述图片') TEXT  default('')" json:"imgs"`
	ConsultantUid int       `xorm:"not null comment('咨询人uid') index UNSIGNED INT" json:"consultant_uid" validate:"required"`
	Status        string    `xorm:"not null default '处理中' comment('处理中、待人工咨询、人工咨询中、已完成') VARCHAR(10)" json:"status"`
	CreateTime    int       `xorm:"not null UNSIGNED INT" json:"create_time"`
	UpdateTime    time.Time `xorm:"not null updated DateTime default(CURRENT_TIMESTAMP)" json:"-"`
}

// “咨询”沟通记录
type ConsultationReply struct {
	Id              int       `xorm:"not null pk autoincr UNSIGNED INT" json:"id"`
	ConsultationId  int       `xorm:"not null comment('咨询id') index UNSIGNED INT" json:"consultation_id"`
	CommunicatorUid int       `xorm:"not null comment('沟通人uid') UNSIGNED INT" json:"communicator_uid"`
	Type            string    `xorm:"not null comment('回复类型，answer,query') VARCHAR(20)" json:"type" validate:"required,oneof=answer query"`
	Content         string    `xorm:"not null comment('回复内容') LONGTEXT" json:"content" validate:"required"`
	CreateTime      int       `xorm:"not null UNSIGNED INT" json:"create_time"`
	UpdateTime      time.Time `xorm:"not null updated DateTime default(CURRENT_TIMESTAMP)" json:"-"`
}

//创建咨询
func (consul *Consultation) Create() error {
	_, err := Db.InsertOne(consul)
	return err
}

//删除咨询
func (consul *Consultation) delete() error {
	if consul.Id == 0 {
		err := errors.New("model:必须指定id值")
		return err
	}
	_, err := Db.Delete(consul)
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

//创建咨询回复记录
func (consul *Consultation) AddReply(record *ConsultationReply) error {
	if consul.Id == 0 {
		err := errors.New("model:必须指定id值")
		return err
	}
	record.ConsultationId = consul.Id
	_, err := Db.InsertOne(record)
	return err
}

//删除咨询回复记录
func (consul *Consultation) deleteReply(record *ConsultationReply) error {
	if consul.Id == 0 {
		err := errors.New("model:必须指定id值")
		return err
	}
	record.ConsultationId = consul.Id
	_, err := Db.Delete(record)
	return err
}

type replyInfo struct {
	Id              int    `json:"id"`
	ConsultationId  int    `json:"consultation_id"`
	CommunicatorUid int    `json:"communicator_uid"`
	Type            string `json:"type"`
	Content         string `json:"content"`
	NickName        string `json:"nick_name"`
	AvatarUrl       string `json:"avatar_url"`
	Phone           string `json:"phone"`
	CreateTime      int    `json:"create_time"`
}

//获取咨询沟通记录表
func (consul *Consultation) ListReply() (replyInfoList []replyInfo, err error) {
	if consul.Id == 0 {
		err = errors.New("model:必须指定id值")
		return
	}
	err = Db.Table("consultation_reply").
		Join("INNER", "user", "user.id = consultation_reply.communicator_uid").
		Where("consultation_id=?", consul.Id).
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
		Find(&replyInfoList)
	return
}

//用户历史咨询记录列表
func ConsultationList(uid int) ([]Consultation, error) {
	consultationList := []Consultation{}
	err := Db.Table("consultation").
		Where("consultation.consultant_uid = ?", uid).
		Desc("create_time").
		Find(&consultationList)
	return consultationList, err
}

type consultationWithUserInfo struct {
	Id            int    `json:"id"`
	Question      string `json:"question"`
	Imgs          string `json:"imgs"`
	Status        string `json:"status"`
	CreateTime    int    `json:"create_time"`
	ConsultantUid int    `json:"consultant_uid"`
	NickName      string `json:"nick_name"`
	AvatarUrl     string `json:"avatar_url"`
	Phone         string `json:"phone"`
}

//获取咨询信息
func ConsultationGetWithUserInfo(consultationId int) (*consultationWithUserInfo, error) {
	consultationInfo := &consultationWithUserInfo{}
	_, err := Db.Table("consultation").
		Join("INNER", "user", "user.id = consultation.consultant_uid").
		Where("consultation.id=?", consultationId).
		Cols(
			"consultation.id",
			"consultation.question",
			"consultation.imgs",
			"consultation.status",
			"consultation.create_time",
			"consultation.consultant_uid",
			"user.nick_name",
			"user.avatar_url",
			"user.phone",
		).
		Get(consultationInfo)
	if consultationInfo.Id == 0 {
		consultationInfo = nil
	}
	return consultationInfo, err
}
