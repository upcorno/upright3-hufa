package service

import (
	"law/model"

	"xorm.io/xorm"
)

type RightsProtectionSearch struct {
	Classification  string `json:"classification" query:"classification"`
	CreateTimeMin   int    `json:"create_time_min" query:"create_time_min"`
	CreateTimeMax   int    `json:"create_time_max" query:"create_time_max"`
	CustomerAddress string `json:"customer_address" query:"customer_address"`
}

type RightsProtectionInfo struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	Phone          string `json:"phone"`
	CreateTime     int    `json:"create_time"`
	Classification string `json:"classification"`
}

//侵权监测列表搜索
func RightsProtectionList(page *model.Page, search *RightsProtectionSearch) (*model.PageResult, error) {
	protectionInfo := []RightsProtectionInfo{}
	sess := model.Db.NewSession()
	sess.Table("rights_protection")
	sess.Join("INNER", "protection_return_visit", "rights_protection.id = protection_return_visit.protection_id")
	sess.Cols(
		"rights_protection.id",
		"rights_protection.name",
		"rights_protection.phone",
		"rights_protection.create_time",
		"protection_return_visit.classification",
	)
	dealProtectionSearch(sess, search)
	pageResult, err := page.GetResults(sess, &protectionInfo)
	if err != nil {
		return nil, err
	}
	return pageResult, err
}

func dealProtectionSearch(sess *xorm.Session, search *RightsProtectionSearch) {
	if search.Classification != "" {
		sess.Where("protection_return_visit.classification = ?", search.Classification)
	}
	if search.CustomerAddress != "" {
		sess.Where("protection_return_visit.customer_address like ?", "%"+search.CustomerAddress+"%")
	}
	if search.CreateTimeMin != 0 {
		sess.Where("rights_protection.create_time > ?", search.CreateTimeMin)
	}
	if search.CreateTimeMax != 0 {
		sess.Where("rights_protection.create_time < ?", search.CreateTimeMax)
	}
}
