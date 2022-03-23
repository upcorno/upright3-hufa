package model

import (
	_ "github.com/go-sql-driver/mysql"
)

//创建咨询记录
func ConsultationRecordCreate(record *ConsultationRecord) error {
	_, err := Db.InsertOne(record)
	return err
}