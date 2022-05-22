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

//创建咨询
func (consul *Consultation) Create() (err error) {
	if consul.Question == "" || consul.ConsultantUid == 0 || consul.Status == "" {
		err = errors.New("Question、ConsultantUid、Status不可以为空值")
		return
	}
	consul.CreateTime = int(time.Now().Unix())
	_, err = Db.InsertOne(consul)
	return err
}

//删除咨询
func (consul *Consultation) delete() error {
	if consul.Id == 0 {
		err := errors.New("model:必须指定id值")
		return err
	}
	_, err := Db.Delete(&Consultation{Id: consul.Id})
	return err
}

func (consul *Consultation) Update(columns ...string) (err error) {
	if consul.Id == 0 {
		err := errors.New("model:必须指定id值")
		return err
	}
	_, err = Db.Cols(columns...).Update(consul, &Consultation{Id: consul.Id})
	return
}

//用户历史咨询记录列表
func ConsultationList(uid int) (consultationList []Consultation, err error) {
	consultationList = []Consultation{}
	err = Db.Table("consultation").
		Where("consultation.consultant_uid = ?", uid).
		Desc("create_time").
		Find(&consultationList)
	return
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
