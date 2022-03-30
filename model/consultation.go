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
func ConsultationGet(consultationId int) (Consultation, error) {
	consultation := Consultation{}
	_, err := Db.Table("consultation").Where("id=?", consultationId).Get(&consultation)
	return consultation, err
}
