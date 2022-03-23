package model

import (
	_ "github.com/go-sql-driver/mysql"
)

func ConsultationCreate(consul *Consultation) error {
	_, err := Db.InsertOne(consul)
	return err
}