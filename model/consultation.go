package model

import (
	_ "github.com/go-sql-driver/mysql"
)

//创建咨询
func ConsultationCreate(consul *Consultation) error {
	_, err := Db.InsertOne(consul)
	return err
}

//设置咨询状态
func ConsultationSetStatus(consultationId int, status string) error {
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
func ConsultationGet(consultationId int) (map[string]string, error) {
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
	return consultationInfo, err
}

//创建咨询记录
func ConsultationAddReply(record *ConsultationReply) error {
	_, err := Db.InsertOne(record)
	return err
}

//获取咨询沟通记录表
func ConsultationListReply(consultationId int) ([]map[string]string, error) {
	recordInfoList := []map[string]string{}
	err := Db.Table("consultation_reply").
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
			"create_time",
		).
		Asc("create_time").
		Find(&recordInfoList)
	return recordInfoList, err
}
