package model

import (
	_ "github.com/go-sql-driver/mysql"
)

//插入回访记录
func ProtectionRetureVisitAdd(protectionReturnVisit *ProtectionReturnVisit) error {
	_, err := Db.InsertOne(protectionReturnVisit)
	return err
}

//更新回访记录
func ProtectionRetureVisitUpdate(protectionReturnVisit *ProtectionReturnVisit) error {
	_, err := Db.Update(protectionReturnVisit, &ProtectionReturnVisit{ProtectionId: protectionReturnVisit.ProtectionId})
	return err
}
