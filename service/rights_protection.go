package service

import (
	"law/model"

	"xorm.io/xorm"
)

type RightsProtectionDealInfo struct {
	Id              int    `json:"id" form:"id" query:"id" validate:"required,gt=0"`
	DealResult      string `json:"deal_result" query:"deal_result" form:"deal_result" validate:"required,oneof=未回访 有合作意向 无合作意向 已合作"`
	CustomerAddress string `json:"customer_address" form:"customer_address" query:"customer_address"`
	DealRemark      string `json:"deal_remark" form:"deal_remark" query:"deal_remark"`
}

func RightsProtectionSetDealInfo(dealInfo *RightsProtectionDealInfo) error {
	protection := &model.RightsProtection{
		DealResult:      dealInfo.DealResult,
		CustomerAddress: dealInfo.CustomerAddress,
		DealRemark:      dealInfo.DealRemark,
	}
	_, err := model.Db.Update(protection, &model.RightsProtection{Id: dealInfo.Id})
	return err
}

type RightsProtectionSearch struct {
	DealResult      string `json:"deal_result" query:"deal_result"`
	DealRemark      string `json:"deal_remark" query:"deal_remark"`
	CustomerAddress string `json:"customer_address" query:"customer_address"`
	CreateTimeMin   int    `json:"create_time_min" query:"create_time_min"`
	CreateTimeMax   int    `json:"create_time_max" query:"create_time_max"`
}

type RightsProtectionInfo struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	CreateTime int    `json:"create_time"`
	DealResult string `json:"deal_result"`
}

//侵权监测列表搜索
func RightsProtectionBackendList(page *model.Page, search *RightsProtectionSearch) (*model.PageResult, error) {
	protectionInfo := []RightsProtectionInfo{}
	sess := model.Db.NewSession()
	sess.Table("rights_protection")
	sess.Cols(
		"id",
		"name",
		"phone",
		"create_time",
		"deal_result",
	)
	dealProtectionSearch(sess, search)
	pageResult, err := page.GetResults(sess, &protectionInfo)
	if err != nil {
		return nil, err
	}
	return pageResult, err
}

func dealProtectionSearch(sess *xorm.Session, search *RightsProtectionSearch) {
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
