package model

import (
	"errors"

	_ "github.com/go-sql-driver/mysql"
)

//维权意向添加
func RightsProtectionAdd(rightsProtection *RightsProtection) error {
	has, err := Db.Exist(&RightsProtection{
		CreatorUid: rightsProtection.CreatorUid,
	})
	if err != nil {
		return err
	}
	if has {
		return errors.New("系统已添加您的维权意向，请勿重复添加！")
	} else {
		_, err = Db.InsertOne(rightsProtection)
		return err
	}
}