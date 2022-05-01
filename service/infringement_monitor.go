package service

import (
	"law/model"

	"xorm.io/xorm"
)

type InfringementMonitorDealInfo struct {
	Id              int    `json:"id" form:"id" query:"id" validate:"required,gt=0"`
	DealResult      string `json:"deal_result" query:"deal_result" form:"deal_result" validate:"required,oneof=未回访 有合作意向 无合作意向 已合作"`
	CustomerAddress string `json:"customer_address" form:"customer_address" query:"customer_address"`
	DealRemark      string `json:"deal_remark" form:"deal_remark" query:"deal_remark"`
}

func InfringementMonitorSetDealInfo(dealInfo *InfringementMonitorDealInfo) error {
	monitor := &model.InfringementMonitor{
		DealResult:      dealInfo.DealResult,
		CustomerAddress: dealInfo.CustomerAddress,
		DealRemark:      dealInfo.DealRemark,
	}
	_, err := model.Db.Update(monitor, &model.InfringementMonitor{Id: dealInfo.Id})
	return err
}

type InfringementMonitorSearch struct {
	DealResult      string `json:"deal_result" query:"deal_result"`
	DealRemark      string `json:"deal_remark" query:"deal_remark"`
	CustomerAddress string `json:"customer_address" query:"customer_address"`
	CreateTimeMin   int    `json:"create_time_min" query:"create_time_min"`
	CreateTimeMax   int    `json:"create_time_max" query:"create_time_max"`
}

type InfringemetMonitorInfo struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	CreateTime int    `json:"create_time"`
	DealResult string `json:"deal_result"`
}

//侵权监测列表搜索
func InfringementMonitorBackendList(page *model.Page, search *InfringementMonitorSearch) (*model.PageResult, error) {
	detectionInfo := []InfringemetMonitorInfo{}
	sess := model.Db.NewSession()
	sess.Table("infringement_monitor")
	sess.Cols(
		"id",
		"name",
		"phone",
		"create_time",
		"deal_result",
	)
	dealDetectionSearch(sess, search)
	pageResult, err := page.GetResults(sess, &detectionInfo)
	if err != nil {
		return nil, err
	}
	return pageResult, err
}

func dealDetectionSearch(sess *xorm.Session, search *InfringementMonitorSearch) {
	if search.DealResult != "" {
		sess.Where("deal_result = ?", search.DealResult)
	}
	if search.CustomerAddress != "" {
		sess.Where("customer_address like ?", "%"+search.CustomerAddress+"%")
	}
	if search.CreateTimeMin != 0 {
		sess.Where("create_time >= ?", search.CreateTimeMin)
	}
	if search.CreateTimeMax != 0 {
		sess.Where("create_time <= ?", search.CreateTimeMax)
	}
}
