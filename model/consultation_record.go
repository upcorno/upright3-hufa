package model

import (
	_ "github.com/go-sql-driver/mysql"
)

//创建咨询记录
func ConsultationRecordCreate(record *ConsultationRecord) error {
	_, err := Db.InsertOne(record)
	return err
}

//获取咨询沟通记录表
func ConsultationRecordList(consultationId int) ([]ConsultationRecord, error) {
	recordList := []ConsultationRecord{}
	err := Db.Table("consultation_record").Where("consultation_id=?", consultationId).Asc("create_time").Find(&recordList)
	return recordList, err
}
