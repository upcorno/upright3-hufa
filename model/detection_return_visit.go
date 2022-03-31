package model

import (
	_ "github.com/go-sql-driver/mysql"
)

//插入回访记录
func DetectionRetureVisitAdd(detectionReturnVisit *DetectionReturnVisit) error {
	_, err := Db.InsertOne(detectionReturnVisit)
	return err
}

//更新回访记录
func DetectionRetureVisitUpdate(detectionReturnVisit *DetectionReturnVisit) error {
	_, err := Db.Update(detectionReturnVisit, &DetectionReturnVisit{DetectionId: detectionReturnVisit.DetectionId})
	return err
}