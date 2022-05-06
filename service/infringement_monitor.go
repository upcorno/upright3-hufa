package service

import (
	"errors"
	"law/model"
	"time"

	"xorm.io/xorm"
)

type InfringementMonitorDealInfo struct {
	Id              int     `json:"id" query:"id" form:"id" validate:"required,gt=0"`
	DealResult      string  `json:"deal_result" query:"deal_result" form:"deal_result" validate:"required,oneof=未回访 有合作意向 无合作意向 已合作"`
	CustomerAddress *string `json:"customer_address" form:"customer_address" query:"customer_address" validate:"min=0,max=50"`
	DealRemark      *string `json:"deal_remark" form:"deal_remark" query:"deal_remark" validate:"required"`
}

func InfringementMonitorSetDealInfo(id int, dealInfo *InfringementMonitorDealInfo) error {
	bean := &model.InfringementMonitor{
		DealResult:      dealInfo.DealResult,
		CustomerAddress: *dealInfo.CustomerAddress,
		DealRemark:      *dealInfo.DealRemark,
	}
	_, err := model.Db.Cols("deal_result", "customer_address", "deal_remark").Update(bean, &model.InfringementMonitor{Id: id})
	return err
}

type InfringementMonitorBaseInfo struct {
	Name         string  `json:"name" form:"name" query:"name" validate:"min=1,max=16"`
	Phone        string  `json:"phone" form:"phone" query:"phone" validate:"min=3,max=20"`
	Organization *string `json:"organization" form:"organization" query:"organization" validate:"min=0,max=60"`
	Description  *string `json:"description" form:"description" query:"description" validate:"min=0"`
	Resume       *string `json:"resume" form:"resume" query:"resume" validate:"min=0"`
}

func InfringementMonitorUpdateBaseInfo(uid int, baseInfo *InfringementMonitorBaseInfo) error {
	bean := &model.InfringementMonitor{
		Name:         baseInfo.Name,
		Phone:        baseInfo.Phone,
		Organization: *baseInfo.Organization,
		Description:  *baseInfo.Description,
		Resume:       *baseInfo.Resume,
	}
	_, err := model.Db.Cols("name", "phone", "organization", "description", "resume").Update(bean, &model.InfringementMonitor{CreatorUid: uid})
	return err
}

func InfringementMonitorAdd(baseInfo *InfringementMonitorBaseInfo, creatorUid int) (id int, err error) {
	has, err := model.Db.Exist(&model.InfringementMonitor{
		CreatorUid: creatorUid,
	})
	if err != nil {
		return
	}
	if has {
		err = errors.New("系统已添加您的侵权监测，请勿重复添加！")
		return
	}
	bean := model.InfringementMonitor{
		Name:         baseInfo.Name,
		Phone:        baseInfo.Phone,
		Organization: *baseInfo.Organization,
		Description:  *baseInfo.Description,
		Resume:       *baseInfo.Resume,
		CreatorUid:   creatorUid,
		CreateTime:   int(time.Now().Unix()),
	}
	bId, err := model.Db.InsertOne(bean)
	id = int(bId)
	return
}

func InfringementMonitorBgGet(beanId int) (model.InfringementMonitor, error) {
	bean := model.InfringementMonitor{}
	_, err := model.Db.Table("infringement_monitor").Where("id=?", beanId).Get(&bean)
	return bean, err
}

func InfringementMonitorGet(creatorUid int) (model.InfringementMonitor, error) {
	bean := model.InfringementMonitor{}
	_, err := model.Db.Table("infringement_monitor").Where("creator_uid=?", creatorUid).Get(&bean)
	return bean, err
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

func InfringementMonitorBackendList(page *model.Page, search *InfringementMonitorSearch) (*model.PageResult, error) {
	searchInfo := []InfringemetMonitorInfo{}
	sess := model.Db.NewSession()
	sess.Table("infringement_monitor")
	sess.Cols(
		"id",
		"name",
		"phone",
		"create_time",
		"deal_result",
	)
	dealInfringementMonitorSearch(sess, search)
	pageResult, err := page.GetResults(sess, &searchInfo)
	if err != nil {
		return nil, err
	}
	return pageResult, err
}

func dealInfringementMonitorSearch(sess *xorm.Session, search *InfringementMonitorSearch) {
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
