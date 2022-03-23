package model

import (
	_ "github.com/go-sql-driver/mysql"
)

//创建咨询
func ConsultationCreate(consul *Consultation) error {
	_, err := Db.InsertOne(consul)
	return err
}