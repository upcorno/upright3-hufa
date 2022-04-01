package service

import (
	"law/model"

	"xorm.io/xorm"
)

type InfringementDetectionSearch struct {
	Classification  string `json:"classification" query:"classification"`
	CreateTimeMin   int    `json:"create_time_min" query:"create_time_min"`
	CreateTimeMax   int    `json:"create_time_max" query:"create_time_max"`
	CustomerAddress string `json:"customer_address" query:"customer_address"`
}

type InfringemetDetectionInfo struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	Phone          string `json:"phone"`
	CreateTime     int    `json:"create_time"`
	Classification string `json:"classification"`
}

//侵权监测列表搜索
func InfringementDetectionList(page *model.Page, search *InfringementDetectionSearch) (*model.PageResult, error) {
	detectionInfo := []InfringemetDetectionInfo{}
	sess := model.Db.NewSession()
	sess.Table("infringement_detection")
	sess.Join("INNER", "detection_return_visit", "infringement_detection.id = detection_return_visit.detection_id")
	sess.Cols(
		"infringement_detection.id",
		"infringement_detection.name",
		"infringement_detection.phone",
		"infringement_detection.create_time",
		"detection_return_visit.classification",
	)
	dealDetectionSearch(sess, search)
	pageResult, err := page.GetResults(sess, &detectionInfo)
	if err != nil {
		return nil, err
	}
	return pageResult, err
}

func dealDetectionSearch(sess *xorm.Session, search *InfringementDetectionSearch) {
	if search.Classification != "" {
		sess.Where("detection_return_visit.classification = ?", search.Classification)
	}
	if search.CustomerAddress != "" {
		sess.Where("detection_return_visit.customer_address like ?", "%"+search.CustomerAddress+"%")
	}
	if search.CreateTimeMin != 0 {
		sess.Where("infringement_detection.create_time > ?", search.CreateTimeMin)
	}
	if search.CreateTimeMax != 0 {
		sess.Where("infringement_detection.create_time < ?", search.CreateTimeMax)
	}
}
