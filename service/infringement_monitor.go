package service

import (
	"errors"
	"law/model"
	"time"

	"xorm.io/xorm"
)

type monitorSrv struct{}

var Monitor = &monitorSrv{}

type InfringementMonitorDealInfo struct {
	Id              int    `json:"id" query:"id" form:"id" validate:"required,gt=0"`
	DealResult      string `json:"deal_result" query:"deal_result" form:"deal_result" validate:"required,oneof=未回访 有合作意向 无合作意向 已合作"`
	CustomerAddress string `json:"customer_address" form:"customer_address" query:"customer_address" validate:"min=0,max=50"`
	DealRemark      string `json:"deal_remark" form:"deal_remark" query:"deal_remark" validate:"min=0"`
}

func (p *monitorSrv) SetDealInfo(id int, dealInfo *InfringementMonitorDealInfo) (err error) {
	bean := &model.InfringementMonitor{
		Id:              id,
		DealResult:      dealInfo.DealResult,
		CustomerAddress: dealInfo.CustomerAddress,
		DealRemark:      dealInfo.DealRemark,
	}
	err = bean.Update(
		"deal_result",
		"customer_address",
		"deal_remark",
	)
	return
}

type InfringementMonitorBaseInfo struct {
	Name         string `json:"name" form:"name" query:"name" validate:"min=1,max=16"`
	Phone        string `json:"phone" form:"phone" query:"phone" validate:"min=3,max=20"`
	Organization string `json:"organization" form:"organization" query:"organization" validate:"min=0,max=60"`
	Description  string `json:"description" form:"description" query:"description" validate:"min=0"`
	Resume       string `json:"resume" form:"resume" query:"resume" validate:"min=0"`
}

func (p *monitorSrv) UpdateBaseInfo(creatorUid int, baseInfo *InfringementMonitorBaseInfo) (err error) {
	bean := &model.InfringementMonitor{
		CreatorUid:   creatorUid,
		Name:         baseInfo.Name,
		Phone:        baseInfo.Phone,
		Organization: baseInfo.Organization,
		Description:  baseInfo.Description,
		Resume:       baseInfo.Resume,
	}
	err = bean.Update("name", "phone", "organization", "description", "resume")
	return
}

func (p *monitorSrv) Add(baseInfo *InfringementMonitorBaseInfo, creatorUid int) (id int, err error) {
	bean := model.InfringementMonitor{CreatorUid: creatorUid}
	has, err := bean.Get()
	if err != nil {
		return
	}
	if has {
		err = errors.New("系统已存在记录，请勿重复添加！")
		return
	}
	bean = model.InfringementMonitor{
		Name:         baseInfo.Name,
		Phone:        baseInfo.Phone,
		Organization: baseInfo.Organization,
		Description:  baseInfo.Description,
		Resume:       baseInfo.Resume,
		CreatorUid:   creatorUid,
		CreateTime:   int(time.Now().Unix()),
	}
	err = bean.Insert()
	id = bean.Id
	return
}

func (p *monitorSrv) BgGet(id int) (bean model.InfringementMonitor, err error) {
	bean.Id = id
	has, err := bean.Get()
	if err != nil {
		return
	}
	if !has {
		err = errors.New("无查询的InfringementMonitor")
		return
	}
	return
}

func (p *monitorSrv) Get(creatorUid int) (bean *model.InfringementMonitor, err error) {
	bean = &model.InfringementMonitor{CreatorUid: creatorUid}
	has, err := bean.Get()
	if err != nil {
		return
	}
	if !has {
		bean = nil
		return
	}
	return
}

type InfringementMonitorSearchParams struct {
	DealResult      string `json:"deal_result" query:"deal_result"`
	DealRemark      string `json:"deal_remark" query:"deal_remark"`
	CustomerAddress string `json:"customer_address" query:"customer_address"`
	CreateTimeMin   int    `json:"create_time_min" query:"create_time_min"`
	CreateTimeMax   int    `json:"create_time_max" query:"create_time_max"`
}

type monitorInfo struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	CreateTime int    `json:"create_time"`
	DealResult string `json:"deal_result"`
}

func (p *monitorSrv) BackendList(page *model.Page, search *InfringementMonitorSearchParams) (*model.PageResult, error) {
	searchInfo := []monitorInfo{}
	sess := model.Db.NewSession()
	sess.Table("infringement_monitor")
	sess.Cols(
		"id",
		"name",
		"phone",
		"create_time",
		"deal_result",
	)
	p.dealSearch(sess, search)
	pageResult, err := page.GetResults(sess, &searchInfo)
	if err != nil {
		return nil, err
	}
	return pageResult, err
}

func (p *monitorSrv) dealSearch(sess *xorm.Session, search *InfringementMonitorSearchParams) {
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
