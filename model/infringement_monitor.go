package model

import (
	"errors"

	_ "github.com/go-sql-driver/mysql"
)

//侵权监测添加
func InfringementMonitorAdd(infringementMonitor *InfringementMonitor) error {
	has, err := Db.Exist(&InfringementMonitor{
		CreatorUid: infringementMonitor.CreatorUid,
	})
	if err != nil {
		return err
	}
	if has {
		return errors.New("系统已添加您的侵权监测，请勿重复添加！")
	} else {
		_, err = Db.InsertOne(infringementMonitor)
		return err
	}
}

//获取侵权监测详情
func InfringementMonitorGet(monitorId int) (InfringementMonitor, error) {
	infringementMonitor := InfringementMonitor{}
	_, err := Db.Table("infringement_monitor").Where("id=?", monitorId).Get(&infringementMonitor)
	return infringementMonitor, err
}
