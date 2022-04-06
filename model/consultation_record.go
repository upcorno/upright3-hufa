package model

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
)
type ConsultationRecordInfo struct {
	Id              int       `json:"id"`
	ConsultationId  int       `json:"consultation_id"`
	CommunicatorUid int       `json:"communicator_uid"`
	Type            string    `json:"type"`
	Content         string    `json:"content"`
	NickName        string    `json:"nick_name"`
	AvatarUrl       string    `json:"avatar_url"`
	Phone           string    `json:"phone"`
	CreateTime      int       `json:"create_time"`
	UpdateTime      time.Time `json:"update_time"`
}

//创建咨询记录
func ConsultationRecordCreate(record *ConsultationRecord) error {
	_, err := Db.InsertOne(record)
	return err
}

//获取咨询沟通记录表
func ConsultationRecordList(consultationId int) ([]ConsultationRecordInfo, error) {
	recordInfoList := []ConsultationRecordInfo{}
	err := Db.Table("consultation_record").
	Join("INNER", "user", "user.id = consultation_record.Communicator_uid").
	Where("consultation_id=?", consultationId).
	Cols(
		"consultation_record.id",
		"consultation_record.consultation_id",
		"consultation_record.communicator_uid",
		"consultation_record.type",
		"consultation_record.content",
		"user.nick_name",
		"user.avatar_url",
		"user.phone",
		"consultation_record.create_time",
		"consultation_record.update_time",
	).
	Asc("create_time").
	Find(&recordInfoList)
	return recordInfoList, err
}
