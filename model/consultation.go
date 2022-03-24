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