package model

import (
	_ "github.com/go-sql-driver/mysql"
)

func DetectionRetureVisitAdd(detectionReturnVisit *DetectionReturnVisit) error {
	_, err := Db.InsertOne(detectionReturnVisit)
	return err
}