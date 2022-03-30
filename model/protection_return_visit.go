package model

import (
	_ "github.com/go-sql-driver/mysql"
)

func ProtectionRetureVisitAdd(protectionReturnVisit *ProtectionReturnVisit) error {
	_, err := Db.InsertOne(protectionReturnVisit)
	return err
}