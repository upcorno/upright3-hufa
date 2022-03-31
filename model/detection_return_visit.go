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

//获取回访记录
func DetectionReturnVisitGet(detectionId int) (DetectionReturnVisit, error) {
	returnVisit := DetectionReturnVisit{}
	_, err := Db.Table("detection_return_visit").Where("detection_id=?", detectionId).Get(&returnVisit)
	return returnVisit, err
}