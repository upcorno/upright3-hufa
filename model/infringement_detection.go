package model

import (
	"errors"

	_ "github.com/go-sql-driver/mysql"
)

//侵权监测添加
func InfringementDetectionAdd(infringementDetection *InfringementDetection) error {
	has, err := Db.Exist(&InfringementDetection{
		CreatorUid: infringementDetection.CreatorUid,
	})
	if err != nil {
		return err
	}
	if has {
		return errors.New("系统已添加您的侵权监测，请勿重复添加！")
	} else {
		_, err = Db.InsertOne(infringementDetection)
		return err
	}
}

//获取侵权监测详情
func InfringementDetectionGet(detectionId int) (InfringementDetection, error) {
	infringementDetection := InfringementDetection{}
	_, err := Db.Table("infringement_detection").Where("id=?", detectionId).Get(&infringementDetection)
	return infringementDetection, err
}
